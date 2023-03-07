package main

import (
	"errors"
	"fmt"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

func explore(cfg *pokeapi.Config, input ...string) error {
	if len(input) == 0 {
		return errors.New("must provide a location area to explore")
	}

	locationArea := input[0]
	resource, err := cfg.GetLocation(locationArea)
	if err != nil {
		return err
	}

	fmt.Println(resource.PokemonEncounters)

	return nil
}
