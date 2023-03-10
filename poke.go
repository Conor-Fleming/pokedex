package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

var commands = map[string]func(*pokeapi.Config, ...string) error{
	"help":    help,
	"map":     mapcommand,
	"mapb":    mapb,
	"catch":   catch,
	"inspect": inspect,
	"explore": explore,
	"pokedex": pokedex,
	"exit":    exit,
	"clear":   clear,
}

func startPoke(cfg *pokeapi.Config) {
	help(cfg)

	for {
		fmt.Print("pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		//checking for empty input
		if scanner.Text() == "" {
			continue
		}
		args := strings.Fields(scanner.Text())

		if val, ok := commands[args[0]]; ok {
			val(cfg, args[1:]...)
			continue
		}
		invalidCommand()
	}
}

func invalidCommand() {
	fmt.Println("Command not found")
}
