package main

import (
	"flag"
	"log"

	"github.com/nikhilbhatia08/taskflow/pkg/common"
	"github.com/nikhilbhatia08/taskflow/pkg/scheduler"
)

var (
	schedulerPort = flag.String("SchedulerPort", ":9000", "Port on which the scheduler serves requests")
)

func main() {
	dbConnectionString := common.GetDbConnectionString()
	scheduleServer := scheduler.NewSchedulerServer(*schedulerPort, dbConnectionString)
	err := scheduleServer.Start()
	if err != nil {
		log.Fatalf("Unable to start the scheduler: %v\n", err)
	}
}