package pokeapi

import (
	"encoding/json"
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

func (c *Config) GetLocation(name string) (Location, error) {
	url := baseURL + "/location-area/" + name

	if v, ok := c.Cache.Get(url); ok {
		resource := Location{}
		err := json.Unmarshal(v, &resource)
		if err != nil {
			log.Print("(from cache)Unmarshal Name: ", err)
			return Location{}, err
		}

		return resource, nil
	}

	res, err := http.Get(url)
	if err != nil {
		log.Print("Calling API: ", err)
		return Location{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print("Read Response Data: ", err)
		return Location{}, err
	}

	//adding response to cache
	c.Cache.Add(url, body)

	resource := Location{}
	err = json.Unmarshal(body, &resource)
	if err != nil {
		log.Print("(New URL) Unmarshal Name : ", err)
		return Location{}, err
	}

	return resource, nil
}
