package wf

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

func ChildWorkflowB(ctx workflow.Context, param string) (string, error) {
	logger := workflow.GetLogger(ctx)

	logger.Info("WorkflowB: Started\n")

	// Execute local activity
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    10 * time.Second,
		ScheduleToCloseTimeout: 10 * time.Second,
	})

	logger.Info("WorkflowB: Launching ServiceB's Activity\n")
	var result1 string
	err := workflow.ExecuteActivity(ctx, ServiceBActivity, param).Get(ctx, &result1)
	if err != nil {
		return "", err
	}
	logger.Info("WorkflowB: ServiceB's Activity finished\n")

	// Exectute external activity
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout:    10 * time.Second,
		ScheduleToCloseTimeout: 10 * time.Second,
		TaskQueue:              "ServiceC",
	})

	logger.Info("WorkflowB: Launching ServiceC's Activity\n")
	var result2 string
	err = workflow.ExecuteActivity(ctx, "ActivityC", param).Get(ctx, &result2)
	if err != nil {
		return "", err
	}
	logger.Info("WorkflowB: ServiceC's Activity finished\n")

	logger.Info("WorkflowB: Finished\n")
	return result2, nil
}

func ServiceBActivity(ctx context.Context, param string) (string, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("ActivityB: Started\n")
	logger.Info("ActivityB: Finished\n")

	var result string
	return result, nil
}
