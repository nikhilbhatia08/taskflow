package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	pb "github.com/nikhilbhatia08/taskflow/generatedproto"
	"google.golang.org/grpc"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/Queues", GetAllQueues)
	r.HandleFunc("/Queues/{QueueName}", GetAllJobsOfQueue)
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://example.com", "http://localhost:3000"}), // Allowed origins
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),                // Allowed methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),               // Allowed headers
	)

	// Wrap the router with the CORS middleware
	handler := corsOptions(r)

	// Start the server
	fmt.Println("Server running at http://localhost:9005")
	http.ListenAndServe(":9005", handler)
}

func GetAllQueues(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	conn, err := grpc.Dial("localhost:9002", grpc.WithInsecure())
	if err != nil {
		http.Error(w, "There was error executing the request", http.StatusBadGateway)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := pb.NewQueueServiceClient(conn)
	resp, err := client.AllQueuesInfo(ctx, &pb.AllQueuesInfoRequest{})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func GetAllJobsOfQueue(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	conn, err := grpc.Dial("localhost:9003", grpc.WithInsecure())
	if err != nil {
		http.Error(w, "There was error executing the request", http.StatusBadGateway)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client := pb.NewJobServiceClient(conn)
	resp, err := client.GetAllJobsOfParticularTaskQueue(ctx, &pb.GetAllJobsOfParticularTaskQueueRequest{
		QueueName: mux.Vars(r)["QueueName"],
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
