package wf

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
)

func ActivityC(ctx context.Context, param string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ActivityC: Started\n")

	logger.Info("ActivityC: About to \"Process\"\n")

	time.Sleep(5 * time.Second)

	logger.Info("ActivityC: Finished\n")
	return param, nil
}
