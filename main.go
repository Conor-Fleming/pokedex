package main

import (
	"fmt"
	"time"

	"github.com/Conor-Fleming/pokedex/internal/pokeapi"
	"github.com/Conor-Fleming/pokedex/internal/pokecache"
)

const (
	BANNER = `
====================================
__________       __          ________                 
\______   \____ |  | __ ____ \______ \   ____ ___  ___
 |     ___/  _ \|  |/ // __ \ |    |  \_/ __ \\  \/  /
 |    |  (  <_> )    <\  ___/ | |__|    \  ___/ >    < 
 |____|   \____/|__|_ \\___  >_______  /\___  >__/\_ \
                     \/    \/        \/     \/      \/
                           
====================================
`
)

func main() {
	fmt.Print(BANNER)
	cfg := &pokeapi.Config{
		Cache:   pokecache.NewCache(time.Minute * 5),
		Pokedex: map[string]pokeapi.Pokemon{},
	}

	startPoke(cfg)
}
