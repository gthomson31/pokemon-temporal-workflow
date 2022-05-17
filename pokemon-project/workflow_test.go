package app

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	// Mock activity implementation
	env.OnActivity(LookupKanto, mock.Anything).Return("https://pokeapi.co/api/v2/pokemon-species/44/", nil)
	env.ExecuteWorkflow(PokemonWorkflow, "gloom")
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var current_pokemon_url string
	require.NoError(t, env.GetWorkflowResult(&current_pokemon_url))
	require.Equal(t, "https://pokeapi.co/api/v2/pokemon-species/44/", current_pokemon_url)
}
