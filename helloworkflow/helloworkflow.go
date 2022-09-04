package helloworkflow

import (
	"context"

	"go.temporal.io/sdk/workflow"
)

func Workflow(ctx workflow.Context, name string) (string, error) {
	return "Temporal! " + name, nil
}

func Activity(ctx context.Context) (string, error) {
	return "", nil
}
