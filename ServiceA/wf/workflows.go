package wf

import (
	"go.temporal.io/sdk/workflow"
)

func StartingWorkflow(ctx workflow.Context, param string) (string, error) {
	logger := workflow.GetLogger(ctx)

	logger.Info("WorkflowA: Started\n")

	ctx = workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
		TaskQueue: "ServiceB",
	})

	// Execute the Activity synchronously (wait for the result before proceeding)
	var result string
	future := workflow.ExecuteChildWorkflow(ctx, "ChildWorkflowServiceB", param)
	if err := future.Get(ctx, &result); err != nil {
		workflow.GetLogger(ctx).Error("SimpleChildWorkflow failed.", "Error", err)
		return "", err
	}
	// Make the results of the Workflow available

	logger.Info("WorkflowA: Finished\n")

	return result, nil
}

/*
func sendActivity(ctx context.Context, param string) (string, error) {
	var activityResult string

	workflowOptions := client.StartWorkflowOptions{
		ID:        "serviceB_workflowID",
		TaskQueue: "Service",
	}
	_, err := workflow.ExecuteWorkflow(context.Background(), workflowOptions, StartingWorkflow, "Temporal")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	return activityResult, nil
}
*/
