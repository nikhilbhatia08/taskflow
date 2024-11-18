package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	pb "github.com/nikhilbhatia08/taskflow/generatedproto"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

func main() {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},         // Allow frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},  // Allowed HTTP methods
		AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Allowed headers
		AllowCredentials: true,                                      // Allow credentials (e.g., cookies, authorization headers)
	})
	http.HandleFunc("/Queues", GetAllQueues)

	handlerWithCORS := corsHandler.Handler(http.DefaultServeMux)

	log.Println("Server is running on port 9005")
	log.Fatal(http.ListenAndServe(":9005", handlerWithCORS))
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
