package main

import "fmt"

func help(cfg *config) error {
	fmt.Println("\n Welcome to Pokedex! ")
	fmt.Println("                             ")
	fmt.Println("These Are the Avaliable commands: ")
	fmt.Println()
	fmt.Println("help   - Show you the Help")
	fmt.Println("exit   - Exits the Go REPL ")
	fmt.Println("map    - view the next page of locations ")
	fmt.Println("mapb   - view the last page of locations ")
	fmt.Println("-----------------------------------------")

	return nil
}
