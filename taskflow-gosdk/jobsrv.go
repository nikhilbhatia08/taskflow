package taskflowgosdk

import (
	"context"
	"log"

	pb "github.com/nikhilbhatia08/taskflow/taskflow-gosdk/generatedproto"
)

// The function takes New job to be added in the task queue and then returns an id of the job
func (sdk *ClientConnection) NewJob(req *CreateJobRequest) (string, error) {
	response, err := sdk.JobServiceClient.CreateJob(context.Background(), &pb.CreateJobRequest{
		QueueName: req.QueueName,
		Payload:   req.Payload,
		Retries:   req.Retries,
	})
	if err != nil {
		log.Printf("There was some error encountered %v", err)
		return "", err
	}

	return response.GetId(), nil
}

// THIS FUNCTION SHOULD BE RE WRITTEN
// The function takes the queuename and returns all the jobs running, executed and blocked in an array
// func (sdk *ClientConnection) GetAllJobsOfTaskQueue(QueueName string) ([]*pb.JobInformation, error) {
// 	response, err := sdk.JobServiceClient.GetAllJobsOfParticularTaskQueue(context.Background(), &pb.GetAllJobsOfParticularTaskQueueRequest{
// 		QueueName: QueueName,
// 	})
// 	if err != nil {
// 		log.Printf("Error Fetching all the jobs of a queue with error %v", err)
// 		return nil, err
// 	}

// 	return response.JobInfo, nil
// }
