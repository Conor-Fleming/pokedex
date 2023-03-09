package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func (c *Config) ListLocation(inputURL *string) (LocationShallow, error) {
	url := baseURL + "/location-area"
	if inputURL != nil {
		url = *inputURL
	}
	//if data for URL exists in cache use that instead of requesting
	if v, ok := c.Cache.Get(url); ok {
		fmt.Println("test")
		resource := LocationShallow{}
		err := json.Unmarshal(v, &resource)
		if err != nil {
			log.Print("Unmarshal Name: ", err)
			return LocationShallow{}, err
		}

		return resource, nil
	}

	res, err := http.Get(url)
	if err != nil {
		log.Print("Calling API: ", err)
		return LocationShallow{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print("Read Response Data: ", err)
		return LocationShallow{}, err
	}

	//adding response to cache
	c.Cache.Add(url, body)

	resource := LocationShallow{}
	err = json.Unmarshal(body, &resource)
	if err != nil {
		log.Print("Unmarshal Name: ", err)
		return LocationShallow{}, err
	}

	return resource, nil
}
