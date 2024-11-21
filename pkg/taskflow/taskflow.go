package taskflow

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	pb "github.com/nikhilbhatia08/taskflow/generatedproto"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
)

// I have to reimplement this fucking thing again

var (
	SUCCESS int32 = 200
	FAILURE int32 = 420
)

type CreateJobRequest struct {
	QueueName string
	Payload   string
	Retries   int32
}

// this worker package should be imported to use the cluster
// This acts as an sdk for golang
// In this sdk file the jobservice and the workerservice functions are together and should be sepparated

type gRPCInfo struct {
	port string
}

type Job struct {
	Id      string
	Payload string
}

type CreateJobResponse struct {
	Id     string
	Status string
}

type ExecutionStatus struct {
	Status int32
}

type WorkerClient struct {
	hostString string
	queueName  string
	handler    func(context.Context, *Job) (*ExecutionStatus, error)
}

func Create(ServerPort string) (*gRPCInfo, error) {
	return &gRPCInfo{
		port: ServerPort,
	}, nil
}

func CreateWorkerClient(hostString string, queueName string, handler func(context.Context, *Job) (*ExecutionStatus, error)) *WorkerClient {
	return &WorkerClient{
		hostString: hostString,
		queueName:  queueName,
		handler:    handler,
	}
}

func (g *gRPCInfo) NewTask(ctx context.Context, req *CreateJobRequest) *CreateJobResponse {
	// log.Println(args...)
	conn, err := grpc.Dial(g.port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the server %v", err)
	}
	defer conn.Close()
	client := pb.NewJobServiceClient(conn)

	resp, _ := client.CreateJob(ctx, &pb.CreateJobRequest{
		QueueName: req.QueueName,
		Payload:   req.Payload,
		Retries:   req.Retries,
	})

	log.Println(resp.Status)
	return &CreateJobResponse{
		Id:     resp.Id,
		Status: resp.Status,
	}
}

func (g *gRPCInfo) Run() error {
	conn, err := grpc.Dial(g.port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the server %v", err)
	}
	defer conn.Close()
	client := pb.NewQueueServiceClient(conn)
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		response, err := client.DequeueTask(ctx, &pb.DequeueTaskRequest{
			QueueName: "default",
		})
		if err != nil {
			log.Printf("Error calling GetStatus: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		log.Printf("Status: %s", response)

		g.hehe()

		// Wait before the next poll

		time.Sleep(1 * time.Second)
	}

	return g.AwaitShutdown()
}

func (g *gRPCInfo) AwaitShutdown() error {
	signs := make(chan os.Signal, 1)
	signal.Notify(signs, unix.SIGTERM, unix.SIGINT)
	<-signs
	return nil
}

func (w *WorkerClient) Run() error {
	conn1, err := grpc.Dial("localhost:9003", grpc.WithInsecure())
	conn, err := grpc.Dial(w.hostString, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the grpc server %v", err)
	}
	defer conn.Close()
	defer conn1.Close()

	client := pb.NewQueueServiceClient(conn)
	client2 := pb.NewJobServiceClient(conn1)
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		response, err := client.DequeueTask(ctx, &pb.DequeueTaskRequest{
			QueueName: w.queueName,
		})
		if err != nil {
			log.Printf("Error calling GetStatus: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		if response.GetId() != "-2" {
			status, err := w.handler(ctx, &Job{
				Id:      response.GetId(),
				Payload: response.GetPayload(),
			})
			if err != nil {
				log.Println(err)
			}

			if status.Status == FAILURE {
				log.Println("Failure")
			} else if status.Status == SUCCESS {
				log.Printf("The id is : %v", response.GetId())
				client2.UpdateJobStatus(ctx, &pb.UpdateJobStatusRequest{
					Id:         response.GetId(),
					StatusCode: SUCCESS,
				})
			}
		}

		time.Sleep(1 * time.Second)
	}

	return w.AwaitShutdown()
}

func (w *WorkerClient) AwaitShutdown() error {
	signs := make(chan os.Signal, 1)
	signal.Notify(signs, unix.SIGTERM, unix.SIGINT)
	<-signs
	return nil
}

func (w *WorkerClient) TriggerReRun(id string) {
	conn, err := grpc.Dial("localhost:9003", grpc.WithInsecure())
	if err != nil {
		log.Printf("Error Connecting to the grpc server %v", err)
	}

	defer conn.Close()
	client := pb.NewJobServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	resp, err := client.TriggerJobReRun(ctx, &pb.TriggerReRunRequest{
		Id: id,
	})

	log.Println(resp)
	log.Println(err)
}

// Given the Id parameter It gets the job status of that id
// func (w *WorkerClient) GetJobStatus

// Given the name of the task queue It returns the number of jobs in a task queue
func (w *WorkerClient) GetJobsOfTaskQueue(queueName string) ([]*pb.JobInformation, error) {
	jobsOfTaskQueue := []*pb.JobInformation{}
	conn, err := grpc.Dial(w.hostString, grpc.WithInsecure())
	if err != nil {
		log.Printf("Unable to connect to the server : %v", err)
		return jobsOfTaskQueue, err
	}
	defer conn.Close()

	client := pb.NewJobServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.GetAllJobsOfParticularTaskQueue(ctx, &pb.GetAllJobsOfParticularTaskQueueRequest{
		QueueName: queueName,
	})

	jobsOfTaskQueue = resp.JobInfo
	return jobsOfTaskQueue, nil
}

func (g *gRPCInfo) hehe() {
	log.Println("dskjfhdsf")
	time.Sleep(5 * time.Second)
}
