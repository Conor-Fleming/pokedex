package main

import (
	"fmt"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

func explore(cfg *pokeapi.Config, input ...string) error {
	fmt.Println(input[0])

	return nil
}
