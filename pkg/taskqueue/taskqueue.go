package taskqueue

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	pb "github.com/nikhilbhatia08/taskflow/generatedproto"
	"google.golang.org/grpc"
)

// The QueueService will start with an inital queue name called `default`

// This is the main queue server
type QueueService struct {
	pb.UnimplementedQueueServiceServer
	queueServicePort     string
	queueServicelistener net.Listener
	grpcServer           *grpc.Server
	queues               map[string]*TaskQueue
	mu                   sync.RWMutex
	ctx                  context.Context
	cancel               context.CancelFunc
	wg                   sync.WaitGroup
}

type Task struct {
	Id      string
	Payload map[string]interface{}
}

// The task queue will contain an Id and payload
// We will get the task from the coordinator
type TaskQueue struct {
	Task chan Task
}

func NewQueueService(port string) *QueueService {
	ctx, cancel := context.WithCancel(context.Background())
	return &QueueService{
		queueServicePort: port,
		queues:           make(map[string]*TaskQueue),
		ctx:              ctx,
		cancel:           cancel,
	}
}

func (q *QueueService) Start() error {
	if err := q.StartgRPCServer(); err != nil {
		log.Printf("The error is : %v", err)
		return err
	}

	go q.FormQueue("default")

	return q.AwaitGracefulShutdown()
}

func (q *QueueService) AwaitGracefulShutdown() error {
	// Set up a channel to listen for shutdown signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // Block until a signal is received

	q.wg.Wait()
	// Log shutdown start
	log.Println("Shutting down server...")

	// Create a context with a timeout for forced shutdown if graceful shutdown takes too long
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Stop gRPC server gracefully
	go func() {
		if q.grpcServer != nil {
			q.grpcServer.GracefulStop()
		}
		q.cancel() // Cancel any running tasks
	}()

	// Wait until timeout or successful shutdown
	select {
	case <-ctx.Done():
		log.Println("Server shutdown complete.")
	case <-time.After(5 * time.Second): // Hard stop after timeout
		log.Println("Forcing server shutdown due to timeout.")
		if q.grpcServer != nil {
			q.grpcServer.Stop()
		}
	}

	os.Exit(1)

	return nil
}

func (q *QueueService) StartgRPCServer() error {
	var err error

	q.queueServicelistener, err = net.Listen("tcp", q.queueServicePort)
	if err != nil {
		log.Fatalf("The error is : %v", err)
		return err
	}

	log.Printf("The server has started at port %v", q.queueServicePort)
	q.grpcServer = grpc.NewServer()
	pb.RegisterQueueServiceServer(q.grpcServer, q)

	go func() {
		if err := q.grpcServer.Serve(q.queueServicelistener); err != nil {
			log.Fatalf("gRPC server failed to start %v", err)
		}
	}()
	return nil
}

func (q *QueueService) CheckIfQueueExists(queueName string) bool {
	_, exists := q.queues[queueName]
	return exists
}

func (q *QueueService) FormQueue(queueName string) string {
	q.mu.Lock()
	defer q.mu.Unlock()

	_, exists := q.queues[queueName]
	if exists == true {
		return "Queue Already exists"
	}

	q.queues[queueName] = &TaskQueue{
		Task: make(chan Task, 1000),
	}
	res := "The queue with the name " + queueName + " has formed"
	return res
}

func (q *QueueService) EnqueueTask(ctx context.Context, req *pb.EnqueueTaskRequest) (*pb.EnqueueTaskResponse, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	_, exists := q.queues[req.GetQueueName()]
	if exists == false {
		return &pb.EnqueueTaskResponse{
			Status: "Queue Name provided is not correct and not present",
		}, nil
	}
	var task Task
	task.Id = req.Id
	if err := json.Unmarshal([]byte(req.GetPayload()), &task.Payload); err != nil {
		log.Printf("Failed to parse JSON :%v", err)
		return &pb.EnqueueTaskResponse{
			Status: "Invalid JSON format",
		}, err
	}
	log.Print(task.Payload)
	log.Printf("The id is : %v", task.Id)
	q.queues[req.GetQueueName()].Task <- task

	return &pb.EnqueueTaskResponse{
		Status: "Success",
	}, nil
}

func (q *QueueService) DequeueTask(ctx context.Context, req *pb.DequeueTaskRequest) (*pb.DequeueTaskResponse, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.CheckIfQueueExists(req.GetQueueName()) == false {
		return &pb.DequeueTaskResponse{
			Id:      "-1",
			Payload: "Queue Does Not Exists",
		}, nil
	}

	if len(q.queues[req.GetQueueName()].Task) == 0 {
		return &pb.DequeueTaskResponse{
			Id:      "-2",
			Payload: "The Queue is empty",
		}, nil
	}

	dqTaskRes := <-q.queues[req.GetQueueName()].Task
	jsonPayload, err := json.Marshal(dqTaskRes.Payload)
	if err != nil {
		log.Println("Problem in marshalling the data")
		return &pb.DequeueTaskResponse{
			Id:      "-1",
			Payload: "Queue Does Not Exists",
		}, err
	}
	return &pb.DequeueTaskResponse{
		Id:      dqTaskRes.Id,
		Payload: string(jsonPayload),
	}, nil
}

// func (q *QueueService) GetAllTaskQueues() (error) {

// }
