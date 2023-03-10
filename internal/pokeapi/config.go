package pokeapi

import "github.com/Conor-Fleming/pokedex/internal/pokecache"

type Config struct {
	Cache   pokecache.Cache
	NextURL *string
	PrevURL *string
	Pokedex map[string]Pokemon
}
