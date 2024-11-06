package common

import (
	"context"
	"log"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetDbConnectionString() string {
	conn := "postgres://nehabhatia@localhost:5433/testdb?sslmode=disable"
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