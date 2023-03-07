package main

import (
	"errors"
	"fmt"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

func catch(cfg *pokeapi.Config, name ...string) error {
	if len(name) == 0 {
		return errors.New("must provide a location area to explore")
	}
	pokemonName := name[0]
	//call catch pokemon func that returns a Pokemon obj and an error
	//if catch returns error then print message and return
	pokemon, err := cfg.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	//do catch logic with using pokemon base experience

	//if the catch func was successful add pokemon to pokedex and print success
	fmt.Printf("%s was caught!", pokemonName)
	cfg.Pokedex[pokemonName] = pokemon

	return nil
}
