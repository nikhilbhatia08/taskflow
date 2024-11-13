package coordinator

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	pb "github.com/nikhilbhatia08/taskflow/generatedproto/"
	"github.com/nikhilbhatia08/taskflow/pkg/common"
	"google.golang.org/grpc"
)

const (
	scanInterval = 10 * time.Second
)

type CoordinatorService struct {
	pb.UnimplementedCoordinatorServiceServer
	serverPort string
	listener net.Listener
	grpcServer *grpc.Server
	dbConnectionString string
	dbpool *pgxpool.Pool
	ctx context.Context
	cancel context.CancelFunc
}

func NewCoordinatorService(port string, dbConnectionString string) *CoordinatorService {
	ctx, cancel := context.WithCancel(context.Background())
	return &CoordinatorService{
		serverPort: port,
		dbConnectionString: dbConnectionString,
		ctx: ctx,
		cancel: cancel,
	}
}

func (s *CoordinatorService) Start() error {
	if err := s.StartgRPCServer(); err != nil {
		return err
	}
	s.dbpool, err := common.ConnectToDatabase(s.ctx, s.dbConnectionString)
	if err != nil {
		return err
	}

	go s.
}

func (s *CoordinatorService) StartgRPCServer() error {
	var err error
	s.listener, err = net.Listen("tcp", s.serverPort)
	if err != nil {
		return err
	}


	log.Printf("Starting Coordinator service on port %v", s.serverPort)
	s.grpcServer = grpc.NewServer()
	pb.RegisterCoordinatorServiceServer(s.grpcServer, s)

	go func() {
		if err := s.grpcServer.Serve(s.listener); err != nil {
			log.Fatalf("gRPC server failed %v", err)
		}
	}()

	return nil
}

func (s *CoordinatorService) ScanDatabase() {
	ticker := time.Ticker(scanInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C: 
			go  s.SendAllScheduledTasksToQueue()
		case <-s.ctx.Done():
			log.Printf("Shutting down database scanner")
			return
		}
	}
} 

func (s *CoordinatorService) SendAllScheduledTasksToQueue() {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	tx, err := s.dbpool.Begin(ctx)
	if err != nil {
		log.Fatalf("Error starting the transactions %v", err)
		return
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil && err.Error() != "tx is closed" {
			log.Printf("ERROR: %#v", err)
			log.Printf("Failed to rollback transaction: %v\n", err)
		}
	}()

	rows, err := tx.Query(ctx, `SELECT id, command FROM tasks WHERE scheduled_at < (NOW() + INTERVAL '30 seconds') AND picked_at IS NULL ORDER BY scheduled_at FOR UPDATE SKIP LOCKED`)

	if err != nil {
		log.Printf("Error Executing the query %v\n", err)
		return 
	}
	defer rows.Close()

	
}