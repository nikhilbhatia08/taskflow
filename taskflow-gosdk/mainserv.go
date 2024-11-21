package taskflowgosdk

import (
	"context"
	"log"
	"os/signal"
	"time"

	"os"

	pb "github.com/nikhilbhatia08/taskflow/taskflow-gosdk/generatedproto"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
)

func NewServer(JobServiceHostString string, QueueServiceHostString string) (*ClientConnection, error) {
	JobServiceConnection, err := grpc.Dial(JobServiceHostString, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error connecting to the Job service : %v", err)
		return &ClientConnection{}, err
	}

	QueueServiceConnection, err := grpc.Dial(QueueServiceHostString, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error connecting to the Job service : %v", err)
		return &ClientConnection{}, err
	}

	return &ClientConnection{
		JobServiceClient:   pb.NewJobServiceClient(JobServiceConnection),
		QueueServiceClient: pb.NewQueueServiceClient(QueueServiceConnection),
	}, nil
}

func (sdk *ClientConnection) Run(runConfig *RunConfigurations) error {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		job, err := sdk.QueueServiceClient.DequeueTask(ctx, &pb.DequeueTaskRequest{
			QueueName: runConfig.QueueName,
		})
		if err != nil {
			log.Printf("There was some error encountered while run : %v", err)
			return err
		}

		if job.Id != QUEUE_IS_EMPTY {
			err := runConfig.handler(context.Background(), &Job{
				Id:      job.GetId(),
				Payload: job.GetPayload(),
			})
			if err != nil {
				sdk.JobServiceClient.UpdateJobStatus(context.Background(), &pb.UpdateJobStatusRequest{
					Id:         job.GetId(),
					StatusCode: int32(JOB_FAILURE),
				})
			} else {
				sdk.JobServiceClient.UpdateJobStatus(context.Background(), &pb.UpdateJobStatusRequest{
					Id:         job.GetId(),
					StatusCode: int32(JOB_SUCCESS),
				})
			}
		}
		time.Sleep(1 * time.Second)
	}

	return sdk.AwaitShutdown()
}

func (sdk *ClientConnection) AwaitShutdown() error {
	signs := make(chan os.Signal, 1)
	signal.Notify(signs, unix.SIGTERM, unix.SIGINT)
	<-signs
	return nil
}
