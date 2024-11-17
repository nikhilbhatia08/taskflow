package workflowservice

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nikhilbhatia08/taskflow/pkg/common"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
)

type WorkflowService struct {
	serverPort string
	listener net.Listener
	grpcServer *grpc.Server
	dbConnectionString string
	dbpool pgxpool.Pool
	ctx context.Context
	cancel context.CancelFunc
}

func NewWorkflowServiceServer(port string, dbConnectionString string) *WorkflowService {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkflowService {
		serverPort: port,
		dbConnectionString: dbConnectionString,
		ctx: ctx,
		cancel: cancel,
	}
}

func (w *WorkflowService) Start() error {
	if err := w.StartgRPCServer(); err != nil {
		log.Println("Error Starting the grpc server %v", err)
		return err
	}
	w.dbpool, err := common.ConnectToDatabase(w.ctx, w.dbConnectionString)
	if err != nil {
		log.Println("Error connecting to the database %v", err)
		return err
	}

	return w.AwaitShutdown()
}

func (w *WorkflowService) StartgRPCServer() error {

}

func (w *WorkflowService) AwaitShutdown() error {
	signs := make(chan os.Signal, 1)
	signal.Notify(signs, unix.SIGTERM, unix.SIGINT)
	<- signs
	return nil
}
