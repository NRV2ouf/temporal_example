package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// Create a new Worker
	worker := worker.New(c, "task-queue-name", worker.Options{})

	// Register Workflows
	worker.RegisterWorkflow(StartingWorkflow)
	// Register Activities
	worker.RegisterActivity(sendActivity)

	workflowOptions := client.StartWorkflowOptions{
		ID:        "serviceA_workflowID",
		TaskQueue: "Service",
	}
	_, err = c.ExecuteWorkflow(context.Background(), workflowOptions, StartingWorkflow, "Temporal")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
}

const YourUpdateName = "	update_name"

func updatedWorkflow(ctx workflow.Context, param string) (string, error) {
	err := workflow.SetUpdateHandler(ctx, YourUpdateName, func(ctx workflow.Context, arg string) (YourUpdateResult, error) {
		fmt.Println("signal reveiced")
		return arg, nil
	})
	if err != nil {
		log.Fatalln("Unable to set update handler", err)
	}

	var result string
	return result, nil
}

func sendActivity(ctx context.Context, param string) (*string, error) {
	// Start Service B

	var result string
	return &result, nil
}
