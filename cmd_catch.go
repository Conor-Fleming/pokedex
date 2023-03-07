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
	// Calculate catch probability based on base experience
	catchProbability := 1.0 / float64(exp)

	// Generate random number between 0 and 1
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Float64()

	// Compare random number to catch probability
	if randomNum <= catchProbability {
		return true // Pokémon caught!
	} else {
		return false // Pokémon escaped :(
	}
}
