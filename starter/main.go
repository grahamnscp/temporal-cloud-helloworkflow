package main

import (
	"context"
	"log"
	"os"

	"github.com/grahamh/temporalcloudhelloworkflow"
	"go.temporal.io/sdk/client"
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

	workflowOptions := client.StartWorkflowOptions{
		ID:        "hello-workflow",
		TaskQueue: "hello-world-taskqueue",
	}

	log.Println("HelloWorkflow starting..")
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, temporalcloudhelloworkflow.HelloWorkflow, "HelloWorkflow test via temporal cloud")
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
