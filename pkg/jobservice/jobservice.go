package jobservice

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	pb "github.com/nikhilbhatia08/taskflow/generatedproto"
	"github.com/nikhilbhatia08/taskflow/pkg/common"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
)

// TODO : Change responses with their designated status codes

// This is a job where id is a unique identifier and
// the payload is a string
type Job struct {
	Id        string
	queueName string
	payload   string
}

type JobService struct {
	pb.UnimplementedJobServiceServer
	serverPort         string
	queueServiceHost   string
	listener           net.Listener
	gRPCServer         *grpc.Server
	dbConnectionString string
	dbpool             *pgxpool.Pool
	ctx                context.Context
	cancel             context.CancelFunc
}

// This gives information about a particular job or a task
type JobInfo struct {
	// this should be filled and will contain information about a job or a task
	Id      string
	Payload string
}

func NewJobServiceServer(port string, dbConnectionString string, queueServiceHost string) *JobService {
	ctx, cancel := context.WithCancel(context.Background())
	return &JobService{
		serverPort:         port,
		dbConnectionString: dbConnectionString,
		queueServiceHost:   queueServiceHost,
		ctx:                ctx,
		cancel:             cancel,
	}
}

func (j *JobService) Start() error {
	var err error
	if err = j.StartgRPCServer(); err != nil {
		log.Printf("Error in the gRPC server %v", err)
		return err
	}

	j.dbpool, err = common.ConnectToDatabase(j.ctx, j.dbConnectionString)
	if err != nil {
		log.Printf("Unable to connect to database")
		return err
	}

	go j.ScanDatabaseAndEnqueueTasks()

	return j.AwaitShutdown()
}

func (j *JobService) StartgRPCServer() error {
	var err error

	j.listener, err = net.Listen("tcp", j.serverPort)
	if err != nil {
		log.Printf("There is an error starting the server : %v", err)
		return err
	}

	log.Printf("Starting the gRPC jobservice server at port %v", j.serverPort)
	j.gRPCServer = grpc.NewServer()
	pb.RegisterJobServiceServer(j.gRPCServer, j)

	go func() {
		if err := j.gRPCServer.Serve(j.listener); err != nil {
			log.Fatalf("gRPC server failed %v", err)
		}
	}()

	return nil
}

func (j *JobService) AwaitShutdown() error {
	signs := make(chan os.Signal, 1)
	signal.Notify(signs, unix.SIGTERM, unix.SIGINT)
	<-signs
	return nil
}

// Creates a job and returns a uuid and status of the job
func (j *JobService) CreateJob(ctx context.Context, req *pb.CreateJobRequest) (*pb.CreateJobResponse, error) {
	tx, err := j.dbpool.Begin(ctx)
	if err != nil {
		return &pb.CreateJobResponse{
			Id:     "",
			Status: "Failed to Create Job",
		}, err
	}

	defer tx.Rollback(ctx)
	queueName := req.GetQueueName()
	payload := req.GetPayload()
	retries := req.GetRetries()

	// TODO : Need to check whether the following queue exists or not
	// TODO : Need to check for the proper values of retries that means it should not be negative and not more that INT_MAX

	uuidString := uuid.New().String()
	_, err = tx.Exec(ctx, `INSERT INTO jobs (id, queuename, payload, retries, maxretries, status) VALUES ($1, $2, $3, $4, $5, $6)`, uuidString, queueName, payload, 0, retries, 1)
	if err != nil {
		log.Printf("The sql statement could not be executed due to an error : %v", err)
		return &pb.CreateJobResponse{
			Id:     "",
			Status: "Failed to Create Job",
		}, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		log.Printf("The sql statement could not be commited : %v", err)
		return &pb.CreateJobResponse{
			Id:     "",
			Status: "Failed to Create Job",
		}, err
	}
	return &pb.CreateJobResponse{
		Id:     uuidString,
		Status: "Success", // this should be changed with their proper status codes
	}, nil
}

func (j *JobService) ScanDatabaseAndEnqueueTasks() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			go j.EnqueueAllReadyJobs()
		case <-j.ctx.Done():
			log.Println("Shutting Down the database scanner")
			return
		}
	}
}

func (j *JobService) EnqueueAllReadyJobs() {
	// log.Println("called")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx, err := j.dbpool.Begin(ctx)
	if err != nil {
		log.Println("Unable to start Transaction")
		return
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil && err.Error() != "tx is closed" {
			log.Printf("Error %#v", err)
			log.Printf("Failed to rollback transaction %v", err)
		}
	}()

	rows, err := tx.Query(ctx, `SELECT id, queuename, payload FROM jobs where status = $1 AND picked_at is NULL`, common.JOB_READY_TO_RUN)
	if err != nil {
		log.Printf("Error querying the jobs %v", err)
	}
	defer rows.Close()

	var jobs []*pb.EnqueueTaskRequest
	for rows.Next() {
		var id, queuename, payload string
		if err := rows.Scan(&id, &queuename, &payload); err != nil {
			log.Printf("Error scanning the rows :%v", err)
			continue
		}
		jobs = append(jobs, &pb.EnqueueTaskRequest{
			Id:        id,
			QueueName: queuename,
			Payload:   payload,
		})
	}

	// log.Println(jobs)

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating the rows %v", err)
		return
	}

	conn, err := grpc.Dial(j.queueServiceHost, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error connecting to the queueService %v", err)
		return
	}

	defer conn.Close()
	client := pb.NewQueueServiceClient(conn)

	for _, job := range jobs {
		// ctx1, cancel1 := context.WithTimeout(context.Background(), 3 * time.Second)
		// defer cancel1()
		res, _ := client.EnqueueTask(context.Background(), job)
		log.Printf("Enquiung Task success %v", res.Status)
		// TODO : the "Success" status should be changed with the proper status code
		if res.Status != "Success" {
			if _, err := tx.Exec(ctx, `UPDATE jobs SET status = $1 WHERE id = $2`, common.BLOCKED_STATE, job.Id); err != nil {
				log.Printf("Failed to update task %v with error : %v", job.Id, err)
				continue
			}
		} else {
			if _, err := tx.Exec(ctx, `UPDATE jobs SET status = $1, picked_at = NOW() WHERE id = $2`, common.JOB_SENT_TO_QUEUE, job.Id); err != nil {
				log.Printf("failed to update task %v with error %v:", job.Id, err)
				continue
			}
		}
	}
	if err := tx.Commit(ctx); err != nil {
		log.Println("Unable to commit transaction")
	}
}

// TODO : These functions should be implemented for the dashboard
// func (j *JobService) GetJobStatusWithId()
// func (j *JobService)
// func (j *JobService) GetAllRunningJobs()
// func (j *JobService) GetAllJobs()
// func (j *JobService) GetAllFailedJobs()
// func (j *JobService) GetAllCompletedJobs()

func (j *JobService) UpdateJobStatus(ctx context.Context, req *pb.UpdateJobStatusRequest) (*pb.UpdateJobStatusResponse, error) {
	ctx1, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx, err := j.dbpool.Begin(ctx1)
	if err != nil {
		log.Printf("Unable to start transaction due to error : %v", err)
		return &pb.UpdateJobStatusResponse{
			StatusCode: common.UPDATE_JOB_STATUS_FAILURE,
		}, err
	}

	defer func() {
		if err := tx.Rollback(ctx1); err != nil && err.Error() != "tx is closed" {
			log.Printf("Unable to create a rollback due to %v", err)
		}
	}()

	var RequestStatusCode int32 = req.GetStatusCode()

	if RequestStatusCode == common.SUCCESS {
		if _, err := tx.Exec(ctx1, `UPDATE jobs SET status = $1, completed_at = NOW() WHERE id = $2`, common.SUCCESS, req.GetId()); err != nil {
			log.Printf("Error Executing the Sql statement with err : %v", err)
			return &pb.UpdateJobStatusResponse{
				StatusCode: common.UPDATE_JOB_STATUS_FAILURE,
			}, err
		}
	} else if RequestStatusCode == common.FAILURE {
		// This particular job Could not be completed due to an error so It should Go into blocked state for greater observability
		if _, err := tx.Exec(ctx1, `UPDATE jobs SET status = $1, failed_at = NOW() WHERE id = $2`, common.BLOCKED_STATE, req.GetId()); err != nil {
			log.Printf("Error Executing the Sql statement with err : %v", err)
			return &pb.UpdateJobStatusResponse{
				StatusCode: common.UPDATE_JOB_STATUS_FAILURE,
			}, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		log.Println("Unable to commit transaction")
	}

	return &pb.UpdateJobStatusResponse{
		StatusCode: common.UPDATE_JOB_STATUS_SUCCESS,
	}, nil
}

// This function is for the dashboard or monitor
func (j *JobService) TriggerJobReRun(ctx context.Context, req *pb.TriggerReRunRequest) (*pb.TriggerReRunResponse, error) {
	ctx1, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx, err := j.dbpool.Begin(ctx1)
	if err != nil {
		log.Printf("Unable to start transaction due to error : %v", err)
		return &pb.TriggerReRunResponse{
			StatusCode: common.UPDATE_JOB_STATUS_FAILURE,
		}, err
	}

	defer func() {
		if err := tx.Rollback(ctx1); err != nil && err.Error() != "tx is closed" {
			log.Printf("Unable to create a rollback due to %v", err)
		}
	}()
	requestId := req.GetId()
	if _, err := tx.Exec(ctx1, `UPDATE jobs SET status = $1, picked_at = NULL WHERE id = $2`, common.JOB_READY_TO_RUN, requestId); err != nil {
		log.Printf("The error is : %v", err)
		return &pb.TriggerReRunResponse{
			StatusCode: common.UPDATE_JOB_STATUS_FAILURE,
		}, err
	}
	if err := tx.Commit(ctx); err != nil {
		log.Println("Unable to commit transaction")
	}

	return &pb.TriggerReRunResponse{
		StatusCode: common.UPDATE_JOB_STATUS_SUCCESS,
	}, nil
}

// This function is to get the particular job details with an Id provided
func (j *JobService) GetJobStatusWithId(ctx context.Context) {

}
