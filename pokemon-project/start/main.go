package main

import (
	"context"
	"fmt"
	"log"
	"pokemon-project/app"

	"go.temporal.io/sdk/client"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	options := client.StartWorkflowOptions{
		ID:        "Pokemon-Workflow",
		TaskQueue: app.PokemonTaskQueue,
	}

	// Pokemon you are requesting details for
	requested_pokemon := "gloom"

	we, err := c.ExecuteWorkflow(context.Background(), options, app.PokemonWorkflow, requested_pokemon)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	// Confirm the Result exists
	var current_pokemon_url string
	err = we.Get(context.Background(), &current_pokemon_url)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	// Return the result back to the UI
	printResults(current_pokemon_url, we.GetID(), we.GetRunID())
}

func printResults(responseData_str string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", responseData_str)
}
