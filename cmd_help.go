package main

import (
	"fmt"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
)

// help displays the helptext
func help(cfg *pokeapi.Config, input ...string) error {
	fmt.Println("\n Welcome to Pokedex! ")
	fmt.Println("                             ")
	fmt.Println("These are the avaliable commands: ")
	fmt.Println()
	fmt.Println("map    			- view the next page of locations")
	fmt.Println("mapb   			- view the last page of locations")
	fmt.Println("explore [area]  	- view list of Pokemon in a given area")
	fmt.Println("catch [name]		- attempt to catch a pokemon with a given name")
	fmt.Println("inspect [name]		- inspect a pokemon from your pokedex")
	fmt.Println("pokedex            	- view the contents of your pokedex")
	fmt.Println()
	fmt.Println("help   			- View the help text")
	fmt.Println("clear  			- clear the screen")
	fmt.Println("exit   			- exit the program")
	fmt.Println("-----------------------------------------")

	return nil
}
