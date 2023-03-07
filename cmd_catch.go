package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

func catch(cfg *pokeapi.Config, name ...string) error {
	if len(name) == 0 {
		return errors.New("must provide a location area to explore")
	}
	pokemonName := name[0]

	//check to see if the given pokemon has already been caught
	if _, ok := cfg.Pokedex[pokemonName]; ok {
		fmt.Printf("You have already caught %s\n", pokemonName)
		return nil
	}
	//call catch pokemon func that returns a Pokemon obj and an error
	//if catch returns error then print message and return
	fmt.Printf("Throwing a Pokeball at %s\n", pokemonName)
	pokemon, err := cfg.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	//if the pokemon escaped, print message and return
	if !catchPokemon(pokemon.BaseExperience) {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	//if pokemon was caught print message and add to pokedex
	fmt.Printf("%s was caught!\n", pokemonName)
	cfg.Pokedex[pokemonName] = pokemon

	return nil
}

func catchPokemon(exp int) bool {
	// Initialize random number generator with current time
	rand.Seed(time.Now().UnixNano())

	// Return true if a random integer between 0 and n/2 is 0
	return rand.Intn(exp/2+1) == 0
}
