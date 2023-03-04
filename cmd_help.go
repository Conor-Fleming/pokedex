package main

import "fmt"

func help(cfg *config) error {
	fmt.Println("\n Welcome to Pokedex! ")
	fmt.Println("                             ")
	fmt.Println("These Are the Avaliable commands: ")
	fmt.Println()
	fmt.Println("help   - Show you the Help")
	fmt.Println("exit   - Exits the Go REPL ")

	return nil
}
