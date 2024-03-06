package wf

import (
	"context"
	"time"
)

func ActivityC(ctx context.Context, param string) (string, error) {
	time.Sleep(5 * time.Second)
	return param, nil
}
