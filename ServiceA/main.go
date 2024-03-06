package main

import (
	"context"
	"log"

	"temporal.io/temporal_example/ServiceA/wf"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// Start the first workflow
	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: "ServiceA",
	}
	_, err = c.ExecuteWorkflow(context.Background(), workflowOptions, wf.StartingWorkflow, "Temporal")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
}
