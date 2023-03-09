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

/*
// get the file descriptor for standard input
	fd := int(os.Stdin.Fd())

	// put the terminal in raw mode to read input byte by byte
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fd, oldState)

	// set the terminal to not echo input back to the user
	if err := term.RawMode(fd); err != nil {
		panic(err)
	}

	// create a new channel to receive input events
	events := make(chan term.Event)

	// start a new goroutine to read input events
	go func() {
		for {
			ev, err := term.ReadEvent(fd)
			if err != nil {
				close(events)
				return
			}
			events <- ev
		}
	}()

	// loop until an interrupt signal is received
	for {
		select {
		case ev := <-events:
			// check if the event is a keyboard event
			if ev.Ch == 0 && ev.Key != 0 {
				switch ev.Key {
				case term.KeyArrowUp:
					fmt.Println("Up arrow pressed")
				case term.KeyArrowDown:
					fmt.Println("Down arrow pressed")
				case term.KeyCtrlC:
					fmt.Println("Interrupt signal received")
					return
				}
			}
		default:
			// do other work here
			time.Sleep(100 * time.Millisecond)
		}
	}
}
*/
