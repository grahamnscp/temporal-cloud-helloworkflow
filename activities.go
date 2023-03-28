package temporalcloudhelloworkflow

import (
	"context"
	"log"

	"go.temporal.io/sdk/activity"
)


func HelloActivity(ctx context.Context, name string) (string, error) {

        logger := activity.GetLogger(ctx)
        logger.Info("HelloActivity", "name", name)

	log.Println("########################### HelloActivity called:", "name", name)

	return "Completed1: " + name, nil
}

func HiActivity(ctx context.Context, name string) (string, error) {

        logger := activity.GetLogger(ctx)
        logger.Info("HiActivity", "name", name)

	log.Println("########################### HiActivity called:", "name", name)

	return "Completed2: " + name, nil
}
