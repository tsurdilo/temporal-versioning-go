package main

import (
	"context"
	"log"
	"temporal/demo/versioning/model"
	"temporal/demo/versioning/workflow"
	"time"

	"go.temporal.io/sdk/client"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	customer := &model.Customer{
		AccountNum:       "c1",
		Name:             "John",
		Email:            "john@john.com",
		CustomerType:     "new",
		DemoWaitDuration: time.Minute,
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:        customer.AccountNum,
		TaskQueue: "versioningGoDemoTaskQueue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflow.CustomerWorkflow, *customer)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// Synchronously wait for the workflow completion.
	var result model.Account
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result.Amount)
}