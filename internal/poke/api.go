package poke

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func callAPI(url *string) (*Resource, error) {
	baseURL := "https://pokeapi.co/api/v2/location-area"
	if url != nil {
		baseURL = *url
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

	resource := &Resource{}
	err = json.Unmarshal(body, &resource)
	if err != nil {
		log.Print("Unmarshal Name: ", err)
		return nil, err
	}

	return resource, nil
}
