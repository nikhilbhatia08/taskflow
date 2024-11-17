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

func (g *gRPCInfo) NewTask(ctx context.Context, req CreateJobRequest) string {
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
	return "recieved"
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
	conn, err := grpc.Dial(w.hostString, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the grpc server %v", err)
	}
	defer conn.Close()

	client := pb.NewQueueServiceClient(conn)
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

		status, err := w.handler(ctx, &Job{
			Id:      response.GetId(),
			Payload: response.GetPayload(),
		})

		if status.Status != 1 {
			log.Println("Failure")
		} else {
			log.Println("Success")
		}

		// Wait before the next poll

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

func (g *gRPCInfo) hehe() {
	log.Println("dskjfhdsf")
	time.Sleep(5 * time.Second)
}
