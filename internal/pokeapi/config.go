package pokeapi

import "github.com/Conor-Fleming/pokedex/internal/pokecache"

type Config struct {
	Cache   pokecache.Cache
	nextURL *string
	prevURL *string
}
