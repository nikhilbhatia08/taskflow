package taskqueue

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/nikhilbhatia08/taskflow/generatedproto"
	"google.golang.org/grpc"
)

// The QueueService will start with an inital queue name called `default`

// This is the main queue server
type QueueService struct {
	pb.UnimplementedQueueServiceServer
	queueServicePort string
	queueServicelistener net.Listener
	grpcServer *grpc.Server
	queues map[string]*TaskQueue
	mu sync.Mutex
	ctx context.Context
	cancel context.CancelFunc
}

type Task struct {
	Id string
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
		queues: make(map[string]*TaskQueue),
		ctx: ctx,
		cancel: cancel,
	}
}

func (q *QueueService) Start() error {
	if err := q.StartgRPCServer(); err != nil {
		return err
	}



	return nil
}

func (q *QueueService) StartgRPCServer() error {
	var err error

	q.queueServicelistener, err = net.Listen("tcp", q.queueServicePort)
	if err != nil {
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

func (q *QueueService) FormQueue(queueName string) string {
	q.mu.Lock()
	defer q.mu.Unlock()

	
}

