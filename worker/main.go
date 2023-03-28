package main

import (
	"log"
        "os"

	"github.com/grahamh/temporalcloudhelloworkflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	clientOptions, err := temporalcloudhelloworkflow.ParseClientOptionFlags(os.Args[1:])
 	if err != nil {
		log.Fatalf("Invalid arguments: %v", err)
	}
        c, err := client.Dial(clientOptions)
	if err != nil {
		log.Fatalln("Unable to make client", err)
	}
	defer c.Close()

	// Start a worker for the taskqueue name
	w := worker.New(c, "hello-world-taskqueue", worker.Options{})

	// Register the workflow and Activities
	w.RegisterWorkflow(temporalcloudhelloworkflow.HelloWorkflow)
	w.RegisterActivity(temporalcloudhelloworkflow.HelloActivity)
	w.RegisterActivity(temporalcloudhelloworkflow.HiActivity)

	// Run the Activity in the Workflow (so it will return to the starter client)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start HelloWorkflow", err)
	}
}
