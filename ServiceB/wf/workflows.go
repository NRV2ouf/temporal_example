package wf

import (
	"context"
	"time"

	"go.temporal.io/sdk/workflow"
)

func ChildWorkflowB(ctx workflow.Context, param string) (string, error) {
	// Execute local activity
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    10 * time.Second,
		ScheduleToCloseTimeout: 10 * time.Second,
	})

	var result1 string
	err := workflow.ExecuteActivity(ctx, ServiceBActivity, param).Get(ctx, &result1)
	if err != nil {
		return "", err
	}

	// Exectute external activity
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    10 * time.Second,
		ScheduleToCloseTimeout: 10 * time.Second,
		TaskQueue:              "ServiceC",
	})
	var result2 string
	err = workflow.ExecuteActivity(ctx, "ActivityC", param).Get(ctx, &result2)
	if err != nil {
		return "", err
	}

	return result2, nil
}

func ServiceBActivity(ctx context.Context, param string) (string, error) {
	var result string
	return result, nil
}
