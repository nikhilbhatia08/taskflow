package scheduler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nikhilbhatia08/taskflow/pkg/common"
)

type SchedulerServer struct {
	port string
	dbConnectionString string
	dbpool *pgxpool.Pool
	ctx context.Context
	cancel context.CancelFunc
	server *http.Server
}

type CommandRequest struct {
	Command string `json:"command"`
	Scheduled_at string `json:"scheduled_at"`
}

type CommandResponse struct {
	Command string `json:"command"`
	Status string `json:"status"`
}

func NewSchedulerServer(port string, dbConnectionString string) *SchedulerServer {
	ctx, cancel := context.WithCancel(context.Background())
	return &SchedulerServer {
		port: port,
		dbConnectionString: dbConnectionString,
		ctx: ctx,
		cancel: cancel,
	}
}

func (s *SchedulerServer) Start() error {
	var err error
	s.dbpool, err = common.ConnectToDatabase(s.ctx, s.dbConnectionString)
	if err != nil {
		return err
	}
	s.server = &http.Server{
		Addr: s.port,
	}

	http.HandleFunc("/submit", s.SubmitHandler)

	go func() {
		err := s.server.ListenAndServe()
		if err != nil {
			log.Fatalf("Unable to start the server: %v\n", err)
		}
	}()

	return s.AwaitGracefulShutdown()
}

func (s *SchedulerServer) SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request CommandRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Println("Received command: ", request.Command)
	log.Println("Scheduled at: ", request.Scheduled_at)

	scheduledTime, err := time.Parse(time.RFC3339, request.Scheduled_at)
	if err != nil {
		http.Error(w, "Invalid scheduled time", http.StatusBadRequest)
		return
	}
	response := struct {
		Command string `json:"command"`
		Status string `json:"status"`
		Scheduled_at int64 `json:"scheduled_at"`
	} {
		Command: request.Command,
		Scheduled_at: scheduledTime.Unix(),
		Status: "Received",
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (s *SchedulerServer) AwaitGracefulShutdown() error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	return s.Stop()
}

func (s *SchedulerServer) Stop() error {
	s.dbpool.Close()
	if s.server != nil {
		ctx, cancel := context.WithTimeout(s.ctx, 5 * time.Second)
		defer cancel()
		return s.server.Shutdown(ctx)
	}
	log.Println("The server has been stopped")
	return nil
}
