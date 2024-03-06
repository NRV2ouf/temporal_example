package main

import (
	"log"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"

	"temporal.io/temporal_example/ServiceB/wf"
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
	w := worker.New(c, "ServiceB", workerOptions)

	// Register child workflow
	registerWOptions := workflow.RegisterOptions{
		Name: "ChildWorkflowServiceB",
	}
	w.RegisterWorkflowWithOptions(wf.ChildWorkflowB, registerWOptions)

	// Register Activity
	w.RegisterActivity(wf.ServiceBActivity)

	// Launch the worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
