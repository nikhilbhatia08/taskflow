# TaskFlow: Task executor and scheduler in Go

![TaskFlow Hero](assets/jobexecution.png)
![TaskFlow Hero](assets/taskflow.png#gh-light-mode-only)
![TaskFlow Hero](assets/taskflowdark.png#gh-dark-mode-only)

TaskFlow is an efficient task executor and scheduler which is used to schedule jobs and tasks and multiple workers can pick those tasks and execute them

## Usage

Import the taskflow-gosdk:
```sh
go get github.com/nikhilbhatia08/taskflow/taskflow-gosdk
```

```go
package tasks

import (
	"context"
	"encoding/json"
	"log"
	"fmt"
	"time"

	taskflowgosdk "github.com/nikhilbhatia08/taskflow/taskflow-gosdk"
)

type EmailPayload struct {
	EmailSenderId string
	EmailRecieverId string
	EmailBody string
}

// Write a function to create a task which consists of queuename, payload and the number of retries
func EmailDelivery() error {
	taskflow, err := taskflowgosdk.NewServer("localhost:9003", "localhost:9002") // The configurations of the jobservice and the queueservice
	if err != nil {
		return err
	}

	payload, err := json.Marshal(EmailPayload{EmailSenderId: "SomeSenderId", EmailRecieverId: "SomeRecieverId", EmailBody: "Some body"})
	if err != nil {
		return err
	}
	taskflow.NewJob(&taskflowgosdk.CreateJobRequest{
		QueueName: "EmailQueue",
		Payload: string(payload),
		Retries: 5,
	})
	return nil
}
```

Create a worker and start executing the jobs:
```go
package tasks

import (
	"context"
	"encoding/json"
	"log"
	"fmt"
	"time"

	taskflowgosdk "github.com/nikhilbhatia08/taskflow/taskflow-gosdk"
)

type EmailPayload struct {
	EmailSenderId string
	EmailRecieverId string
	EmailBody string
}

// Write a function to create a worker to poll to the task queue and execute the jobs
func EmailWorker() error {
	worker, err := taskflowgosdk.NewServer("localhost:9003", "localhost:9002") // The configurations of the jobservice and the queueservice
	if err != nil {
		return err
	}

	worker.Run(&taskflowgosdk.RunConfigurations{
		QueueName: "EmailQueue",
		Handler: EmailDeliveryHandler,
	})
	return nil
}

func EmailDeliveryHandler(ctx context.Context, emailJob *taskflowgosdk.Job) error {
	// Write the job handler logic
}
```

## System Components

## Life of a schedule

## Directory structure

Here's a brief overview of the project's directory structure:

- [`cmd/`](./cmd/): Contains the main entry points for the scheduler, coordinator, task queue and worker services.
- [`pkg/`](./pkg/): Contains the core logic for the scheduler, coordinator, task queue and worker services.
- [`data/`](./data/): Contains SQL scripts to initialize the db.
- [`tests/`](./tests/): Contains integration tests.
- [`*-dockerfile`](./docker-compose.yml): Dockerfiles for building the scheduler, coordinator, task queue and worker services.
- [`docker-compose.yml`](./docker-compose.yml): Docker Compose configuration file for spinning up the entire cluster.
