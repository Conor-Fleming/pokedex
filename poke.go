package main

import (
	"fmt"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

var commands = map[string]func(*pokeapi.Config) error{
	"help": help,
	"map":  mapcommand,
	"mapb": mapb,
	"exit": exit,
}

func startPoke(cfg *pokeapi.Config) {
	help(cfg)

	for {
		fmt.Print("pokedex > ")
		var input string
		fmt.Scanf("%v", &input)
		if value, ok := commands[input]; ok {
			value(cfg)
			continue
		}
		invalidCommand()
	}
}

func invalidCommand() {
	fmt.Println("Command not found")
}
