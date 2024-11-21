package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/nikhilbhatia08/taskflow/pkg/taskflow"
)

// This is an example file to test wheter everything works fine or not

type Task struct {
	Id        string
	queueName string
	payload   string
}

type pay struct {
	id   string
	name string
}

type Request struct {
	id        string
	queueName string
	payload   string
}

type Job struct {
	Id      string
	Payload map[string]interface{}
}

func main() {
	taskflow1, err := taskflow.Create("localhost:9003")
	// The place where the user has to dial
	// The queue that he should dial
	//
	if err != nil {
		log.Fatalf("The error : %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	jsonData := `{
		"name": "Nikhil Bhatia",
		"age": 15,
		"hobbies": {
			"Tennis": 1,
			"football": 2
		}
	}`
	for i := 0; i < 5000; i++ {
		task := taskflow.CreateJobRequest{
			QueueName: "default",
			Payload:   string(jsonData),
			Retries:   5,
		}
		res := taskflow1.NewTask(ctx, &task)
		log.Println(res)
	}

	// worker := taskflow.CreateWorkerClient("localhost:9002", "default", ExecuteJob)
	// // worker.TriggerReRun("1c6becf1-9a86-4908-b6fd-e0e40bb0f736")
	// worker.Run()
}

func ExecuteJob(ctx context.Context, job *taskflow.Job) (*taskflow.ExecutionStatus, error) {
	log.Println(job)
	hehe := job.Id
	log.Println(hehe)
	var haha map[string]interface{}

	err := json.Unmarshal([]byte(job.Payload), &haha)
	if err != nil {
		log.Printf("The error is %v", err)
	}
	log.Println(haha)
	// time.Sleep(5 * time.Second)
	// var j Job
	// err := json.Unmarshal([]byte(job), &j)
	return &taskflow.ExecutionStatus{
		Status: taskflow.SUCCESS,
	}, nil
}
