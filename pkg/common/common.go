package common

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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
	// conn := "postgres://nehabhatia@localhost:5432/testdb?sslmode=disable"
	var missingEnvVars []string

		checkEnvVar := func(envVar, envVarName string) {
			if envVar == "" {
				missingEnvVars = append(missingEnvVars, envVarName)
			}
		}

		dbUser := os.Getenv("POSTGRES_USER")
		checkEnvVar(dbUser, "POSTGRES_USER")

		dbPassword := os.Getenv("POSTGRES_PASSWORD")
		checkEnvVar(dbPassword, "POSTGRES_PASSWORD")

		dbName := os.Getenv("POSTGRES_DB")
		checkEnvVar(dbName, "POSTGRES_DB")

		dbHost := os.Getenv("POSTGRES_HOST")
		if dbHost == "" {
			dbHost = "localhost"
		}

		if len(missingEnvVars) > 0 {
			log.Fatalf("The following required environment variables are not set: %s",
				strings.Join(missingEnvVars, ", "))
		}
	// return conn
	return fmt.Sprintf("postgres://%s:%s@%s:5432/%s", dbUser, dbPassword, dbHost, dbName)
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
