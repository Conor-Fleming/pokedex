package main

import (
	"errors"
	"fmt"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

func mapcommand(cfg *pokeapi.Config, input ...string) error {
	resource, err := cfg.CallAPI(cfg.NextURL)
	if err != nil {
		return err
	}

	cfg.NextURL = resource.Next
	cfg.PrevURL = resource.Previous

	printResults(resource)

	return nil
}

func mapb(cfg *pokeapi.Config, input ...string) error {
	if cfg.PrevURL == nil {
		fmt.Println("Error: You are on the first page")
		return errors.New("No previous page")
	}

	resource, err := cfg.CallAPI(cfg.PrevURL)
	if err != nil {
		return err
	}

	cfg.NextURL = resource.Next
	cfg.PrevURL = resource.Previous

	printResults(resource)

	return nil
}

func printResults(names *pokeapi.Resource) {
	for _, v := range names.Results {
		fmt.Println(v.Name)
	}
}
