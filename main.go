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

func main() {
	fmt.Print(BANNER)
	help()

	commands := map[string]interface{}{
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
			value.(func())()
			continue
		}
		invalidCommand()
	}
}

func invalidCommand() {
	fmt.Println("Command not found")
}

func help() {
	fmt.Println("\n Welcome to Pokedex! ")
	fmt.Println("                             ")
	fmt.Println("These Are the Avaliable commands: ")
	fmt.Println()
	fmt.Println("help   - Show you the Help")
	fmt.Println("exit   - Exits the Go REPL ")
}

func exit() {
	os.Exit(0)
}
