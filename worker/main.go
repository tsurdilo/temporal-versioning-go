package main

import (
	"log"
	"temporal/demo/versioning/activities"
	"temporal/demo/versioning/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "versioningGoDemoTaskQueue", worker.Options{})

	w.RegisterWorkflow(workflow.CustomerWorkflow)
	w.RegisterActivity(activities.CheckCustomerAccount)
	w.RegisterActivity(activities.GetCustomerAccount)
	w.RegisterActivity(activities.UpdateCustomerAccount)
	w.RegisterActivity(activities.SendBonusEmail)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}