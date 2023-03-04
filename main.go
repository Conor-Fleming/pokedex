package main

import (
	"fmt"
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
	cfg := &config{}

	startPoke(cfg)
}

func invalidCommand() {
	fmt.Println("Command not found")
}
