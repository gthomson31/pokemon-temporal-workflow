package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// Workflows are functions that organize Activity method calls.
// Our Workflow will orchestrate the call of a single Activity function.

func PokemonWorkflow(ctx workflow.Context, requested_pokemon string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// declaring the result variable
	var result string

	// Execute the LookupKanto function
	// Return the result as a string.
	err := workflow.ExecuteActivity(ctx, LookupKanto, requested_pokemon).Get(ctx, &result)
	return result, err
}
