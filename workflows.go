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
	_ = UpcertSearchAttribute(ctx, "CustomStringField", "ACTIVITY1")
	var result string
	err := workflow.ExecuteActivity(ctx, HelloActivity, name).Get(ctx, &result)
	if err != nil {
		_ = UpcertSearchAttribute(ctx, "CustomStringField", "FAILED1")
		logger.Error("HelloActivity failed.", "Error", err)
	}

	// sleep (seconds)
	var delay int = 10
	workflow.Sleep(ctx, time.Duration(delay)*time.Second)

	// Run Second Activity
	_ = UpcertSearchAttribute(ctx, "CustomStringField", "ACTIVITY2")
	err = workflow.ExecuteActivity(ctx, HiActivity, name).Get(ctx, &result)
	if err != nil {
		_ = UpcertSearchAttribute(ctx, "CustomStringField", "FAILED2")
		logger.Error("HiActivity failed.", "Error", err)
	}

	_ = UpcertSearchAttribute(ctx, "CustomStringField", "COMPLETE")
	logger.Info("HelloWorkflow workflow completed.", "result", result)

	return result, nil
}

