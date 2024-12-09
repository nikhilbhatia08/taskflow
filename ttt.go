package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nikhilbhatia08/taskflow/generatedproto"
	"google.golang.org/grpc"
)

func main() {
	startTime := time.Now()
	conn, err := grpc.Dial("localhost:9003", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	client := pb.NewJobServiceClient(conn)
	res, err := client.GetAllJobsOfParticularTaskQueue(context.Background(), &pb.GetAllJobsOfParticularTaskQueueRequest{
		QueueName: "default",
		Page:      2,
		PageSize:  1000,
	})
	log.Println(res)
	elapsedTime := time.Since(startTime)
	log.Printf("Time taken : %v", elapsedTime)
}
