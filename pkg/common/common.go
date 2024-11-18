package common

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// these status codes should be uniform throughout the framework
var (
	// Represents the job is ready to be queued and ready to be picked by a worker
	JOB_READY_TO_RUN int32 = 1
	// Represents that a job has been successfully executed
	JOB_SENT_TO_QUEUE int32 = 2
	// Represents blocked state for a job
	BLOCKED_STATE int32 = 3
	// Represents success code for any task, job or process throughout the framework
	SUCCESS int32 = 200
	// Represents success code for update operation of a job
	UPDATE_JOB_STATUS_SUCCESS int32 = 201
	// Represents failure code for any task, job or process throughout the framework
	FAILURE int32 = 420
	// Represents success code for update operation of a job
	UPDATE_JOB_STATUS_FAILURE int32 = 201
)

func GetDbConnectionString() string {
	conn := "postgres://nehabhatia@localhost:5432/testdb?sslmode=disable"
	return conn
}

func ConnectToDatabase(ctx context.Context, conn string) (*pgxpool.Pool, error) {
	var dbpool *pgxpool.Pool
	var err error
	var retryCount int32 = 0
	for retryCount < 5 {
		dbpool, err = pgxpool.Connect(ctx, conn)
		if err != nil {
			retryCount++
		} else {
			break
		}
	}
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	log.Println("Connected to the database")
	return dbpool, nil
}
