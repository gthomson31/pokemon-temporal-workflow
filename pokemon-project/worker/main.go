package main

import (
	"log"
	"pokemon-project/app"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	// This worker hosts both Workflow and Activity functions
	// workflow.go & activity.go

	w := worker.New(c, app.PokemonTaskQueue, worker.Options{})
	w.RegisterWorkflow(app.PokemonWorkflow)
	w.RegisterActivity(app.LookupKanto)
	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
