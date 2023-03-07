package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Config) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	res, err := http.Get(url)
	if err != nil {
		log.Print("Calling API: ", err)
		return Pokemon{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print("Read Response Data: ", err)
		return Pokemon{}, err
	}

	//adding response to cache
	c.Cache.Add(url, body)

	resource := Pokemon{}
	err = json.Unmarshal(body, &resource)
	if err != nil {
		log.Print("(New URL) Unmarshal Name : ", err)
		return Pokemon{}, err
	}

	return resource, nil
}
