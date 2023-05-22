package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

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

	workflowOptions := client.StartWorkflowOptions{
		ID:        "hello-workflow",
		TaskQueue: "hello-world-taskqueue",
	}

	log.Println("HelloWorkflow starting..")
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, hellowf.HelloWorkflow, "HelloWorkflow test via temporal cloud")
	if err != nil {
		log.Fatalln("Unable to execute HelloWorkflow", err)
	}

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get HelloWorkflow result", err)
	}

	log.Println("HelloWorkflow result:", result)
}
