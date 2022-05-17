package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Pokemon_entries []struct {
		Entry_number    int `json:"entry_number"`
		Pokemon_species struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}
	} `json:"pokemon_entries"`
}

// An Activity is just a function that contains your business logic.
// Ours will simply format some text and return it. -- Add Hello to the value of Name
// return the greeting.

func LookupKanto(requested_pokemon string) (string, error) {

	// Get request
	resp, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Println("No response from request")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	// fmt.Println(string(body))
	// Create a Variable with Type Response ( Struct )
	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	var current_pokemon_url string
	var current_pokemon_name string

	// Loop through the data node for the FirstName
	for _, rec := range result.Pokemon_entries {
		current_pokemon_name = rec.Pokemon_species.Name

		if current_pokemon_name == requested_pokemon {
			fmt.Println("Match Found for : " + requested_pokemon)
			current_pokemon_url = rec.Pokemon_species.URL
			break
		}

	}

	return current_pokemon_url, nil
}
