package main

import (
	"log"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporal.io/temporal_example/ServiceC/wf"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// Create a new Worker
	workerOptions := worker.Options{
		WorkerStopTimeout:           5 * time.Minute,
		DisableRegistrationAliasing: false,
	}
	w := worker.New(c, "ServiceC", workerOptions)

	// Register Activity
	registerAOptions := activity.RegisterOptions{
		Name: "ActivityC",
	}
	w.RegisterActivityWithOptions(wf.ActivityC, registerAOptions)

	// Launch the worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
