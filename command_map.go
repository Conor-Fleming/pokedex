package main

import (
	"errors"
	"fmt"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

func mapcommand(cfg *config) error {
	resource, err := pokeapi.CallAPI(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = resource.Next
	cfg.prevURL = resource.Previous

	printResults(resource)

	return nil
}

func mapb(cfg *config) error {
	if cfg.prevURL == nil {
		fmt.Println("Error: You are on the first page")
		return errors.New("No previous page")
	}

	resource, err := pokeapi.CallAPI(cfg.prevURL)
	if err != nil {
		return err
	}

	cfg.nextURL = resource.Next
	cfg.prevURL = resource.Previous

	printResults(resource)

	return nil
}

func printResults(names *pokeapi.Resource) {
	for _, v := range names.Results {
		fmt.Println(v.Name)
	}
}
