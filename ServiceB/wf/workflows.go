package wf

import (
	"context"
	"time"

	"go.temporal.io/sdk/workflow"
)

func ChildWorkflowB(ctx workflow.Context, param string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    10 * time.Second,
		ScheduleToCloseTimeout: 10 * time.Second,
	})

	// Execute the Activity synchronously (wait for the result before proceeding)
	var result string
	err := workflow.ExecuteActivity(ctx, ServiceBActivity, param).Get(ctx, &result)
	if err != nil {
		return "", err
	}
	// Make the results of the Workflow available
	return result, nil
}

func ServiceBActivity(ctx context.Context, param string) (string, error) {
	var result string
	return result, nil
}
