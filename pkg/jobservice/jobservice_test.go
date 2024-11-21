package jobservice

import (
	"context"
	"testing"
	"time"

	"github.com/nikhilbhatia08/taskflow/pkg/taskflow"
)

// TODO : While running the tests start the cluster automatically
// These tests should be run only when we spin up the whole cluster or else we will not be able to properly test everything

func JobCreationTest(t *testing.T) {
	taskflowJobCreator, err := taskflow.Create("localhost:9003")
	if err != nil {
		t.Fatalf("There was no error expected but got error %v", err)
	}

	createdIds := []string{}

	jsonPayload := `{"msg": "Some Payload"}`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for i := 0; i < 5; i++ {
		task := &taskflow.CreateJobRequest{
			QueueName: "default",
			Payload:   jsonPayload,
			Retries:   5,
		}
		resp := taskflowJobCreator.NewTask(ctx, task)
		createdIds = append(createdIds, resp.Id)
	}

}
