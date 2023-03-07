package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

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
