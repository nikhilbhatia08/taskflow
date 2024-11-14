package main

import "github.com/nikhilbhatia08/taskflow/pkg/taskqueue"

func main() {
	taskqueue := taskqueue.NewQueueService(":9002")
	taskqueue.Start()
}
