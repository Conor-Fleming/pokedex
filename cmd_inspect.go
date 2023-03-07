package main

import (
	"errors"
	"fmt"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

func inspect(cfg *pokeapi.Config, name ...string) error {
	if len(name) == 0 {
		return errors.New("must provide a pokemon name to inspect")
	}

	pokemon := name[0]

	//check for existence in map, if found print information
	if val, ok := cfg.Pokedex[pokemon]; ok {
		listInfo(val)
		return nil
	}

	//if the pokemon doesnt exist in map, print error message and return
	fmt.Printf("You havent caught %s yet!\n", pokemon)
	return nil
}

func listInfo(pokemon pokeapi.Pokemon) {
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, val := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Println("Types:")
	for _, val := range pokemon.Types {
		fmt.Printf("  - %s\n", val.Type.Name)
	}
}
