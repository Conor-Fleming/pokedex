package main

import (
	"os"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

func exit(cfg *pokeapi.Config) error {
	os.Exit(0)
	return nil
}
