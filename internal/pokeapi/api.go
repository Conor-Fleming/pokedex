package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Config) CallAPI(url *string) (*Resource, error) {
	baseURL := "https://pokeapi.co/api/v2/location-area"
	if url != nil {
		baseURL = *url
	}
	//if data for URL exists in cache use that instead of requesting
	if v, ok := c.Cache.Get(baseURL); ok {
		resource := &Resource{}
		err := json.Unmarshal(v, &resource)
		if err != nil {
			log.Print("Unmarshal Name: ", err)
			return nil, err
		}

		return resource, nil
	}

	res, err := http.Get(baseURL)
	if err != nil {
		log.Print("Calling API: ", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print("Read Response Data: ", err)
		return nil, err
	}

	//adding response to cache
	c.Cache.Add(baseURL, body)

	resource := &Resource{}
	err = json.Unmarshal(body, &resource)
	if err != nil {
		log.Print("Unmarshal Name: ", err)
		return nil, err
	}

	return resource, nil
}
