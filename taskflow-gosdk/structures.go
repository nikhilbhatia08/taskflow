package taskflowgosdk

import (
	"context"

	pb "github.com/nikhilbhatia08/taskflow/taskflow-gosdk/generatedproto"
)

type ClientConnection struct {
	JobServiceClient   pb.JobServiceClient
	QueueServiceClient pb.QueueServiceClient
}

type Job struct {
	Id      string
	Payload string
}

type RunConfigurations struct {
	QueueName string
	handler   func(context.Context, *Job) error
}

type CreateJobRequest struct {
	QueueName string
	Payload   string
	Retries   int32
}
