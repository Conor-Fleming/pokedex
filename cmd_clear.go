package main

import (
	"os"
	"os/exec"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

func clear(cfg *pokeapi.Config) error {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	return nil
}
