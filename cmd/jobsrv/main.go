package main

import (
	"github.com/nikhilbhatia08/taskflow/pkg/common"
	"github.com/nikhilbhatia08/taskflow/pkg/jobservice"
)

// This is the main entry point of the job service

func main() {
	jobService := jobservice.NewJobServiceServer(":9003", common.GetDbConnectionString(), "localhost:9002")
	jobService.Start()
}
