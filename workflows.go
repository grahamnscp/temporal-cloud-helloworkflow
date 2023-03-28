package temporalcloudhelloworkflow

import (
	"time"

	"go.temporal.io/sdk/workflow"
)


func HelloWorkflow(ctx workflow.Context, name string) (string, error) {

	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("HelloWorkflow workflow started", "name", name)

	// Run first Activity
	var result string
	err := workflow.ExecuteActivity(ctx, HelloActivity, name).Get(ctx, &result)
	if err != nil {
		logger.Error("HelloActivity failed.", "Error", err)
	}

	// Run Second Activity
	err = workflow.ExecuteActivity(ctx, HiActivity, name).Get(ctx, &result)
	if err != nil {
		logger.Error("HiActivity failed.", "Error", err)
	}

	logger.Info("HelloWorkflow workflow completed.", "result", result)

	return result, nil
}

