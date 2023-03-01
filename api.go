package main

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go"
)

func mapcommand() {
	l, err := pokeapi.Resource("location-area")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(l)
}

func mapb() {
	return
}
