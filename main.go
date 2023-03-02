package main

import (
	"fmt"
	"os"
)

const (
	BANNER = `
====================================
   _________________________                
  / __  / __  /| |/ /  /___/
 / /_/_/ /_/ / |   \  / __/ 
/_/   /_____/  |_|\_\/___/
                           
====================================
`
)

type config struct {
	nextURL *string
	prevURL *string
}

func main() {
	fmt.Print(BANNER)

	cfg := &config{}
	help(cfg)

	commands := map[string]func(*config) error{
		"help": help,
		"map":  mapcommand,
		"mapb": mapb,
		"exit": exit,
	}

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

func help(cfg *config) error {
	fmt.Println("\n Welcome to Pokedex! ")
	fmt.Println("                             ")
	fmt.Println("These Are the Avaliable commands: ")
	fmt.Println()
	fmt.Println("help   - Show you the Help")
	fmt.Println("exit   - Exits the Go REPL ")

	return nil
}

func exit(cfg *config) error {
	os.Exit(0)
	return nil
}
