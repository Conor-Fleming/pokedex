package main

import (
	"errors"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

func inspect(cfg *pokeapi.Config, name ...string) error {
	if len(name) == 0 {
		return errors.New("must provide a pokemon name to inspect")
	}

	return nil
}
