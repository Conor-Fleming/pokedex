package main

import (
	"os"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

func exit(cfg *pokeapi.Config, input ...string) error {
	os.Exit(0)

	return nil
}
