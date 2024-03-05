package main

import (
	"context"
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

func StartingWorkflow(ctx workflow.Context, param string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{})

	// Execute the Activity synchronously (wait for the result before proceeding)
	var result string
	err := workflow.ExecuteActivity(ctx, sendActivity, param).Get(ctx, &result)
	if err != nil {
		return "", err
	}
	// Make the results of the Workflow available
	return result, nil
}

func sendActivity(ctx context.Context, param string) (*string, error) {

	updateHandle, err := ctx.client.UpdateWorkflow(context.Background(), updates.YourUpdateWFID, "", updates.YourUpdateName, updateArg)
	if err != nil {
		log.Fatalln("Error issuing Update request", err)
	}

	var result string
	return &result, nil
}
