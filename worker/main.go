package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	hellowf "temporalcloudhelloworkflow"
)

func main() {

	clientOptions, err := hellowf.LoadClientOption()
	if err != nil {
		log.Fatalf("Failed to load Temporal Cloud environment: %v", err)
        }
        c, err := client.Dial(clientOptions)
	if err != nil {
		log.Fatalln("Unable to make client", err)
	}
	defer c.Close()

	// Start a worker for the taskqueue name
	w := worker.New(c, "hello-world-taskqueue", worker.Options{})

	// Register the workflow and Activities
	w.RegisterWorkflow(hellowf.HelloWorkflow)
	w.RegisterActivity(hellowf.HelloActivity)
	w.RegisterActivity(hellowf.HiActivity)

	// Run the Activity in the Workflow (so it will return to the starter client)
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start HelloWorkflow", err)
	}
}
