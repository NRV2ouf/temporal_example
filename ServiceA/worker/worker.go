package main

import (
	"log"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporal.io/temporal_example/ServiceA/wf"
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
	w := worker.New(c, "ServiceA", workerOptions)

	// Register Workflows
	w.RegisterWorkflow(wf.StartingWorkflow)

	// Launch the worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
