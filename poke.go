package main

import "fmt"

type config struct {
	nextURL *string
	prevURL *string
}

var commands = map[string]func(*config) error{
	"help": help,
	"map":  mapcommand,
	"mapb": mapb,
	"exit": exit,
}

func startPoke(cfg *config) {
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
