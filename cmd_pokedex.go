package main

import (
	"fmt"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

func pokedex(cfg *pokeapi.Config, input ...string) error {
	//if empty print message and return
	if len(cfg.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty")
		return nil
	}

	//range over map and print name value for each
	fmt.Println("Your Pokedex:")
	for _, v := range cfg.Pokedex {
		fmt.Printf("  - %s", v.Name)
	}

	return nil
}
