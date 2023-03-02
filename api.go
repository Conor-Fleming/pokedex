package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Resource struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func mapcommand(cfg *config) error {
	resource, err := callAPI(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = resource.Next
	cfg.prevURL = resource.Previous

	printResults(resource)

	return nil
}

func mapb(cfg *config) error {
	if cfg.prevURL == nil {
		fmt.Println("Error: You are on the first page")
		return errors.New("No previous page")
	}

	resource, err := callAPI(cfg.prevURL)
	if err != nil {
		return err
	}

	cfg.nextURL = resource.Next
	cfg.prevURL = resource.Previous

	printResults(resource)

	return nil
}

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

func printResults(names *Resource) {
	for _, v := range names.Results {
		fmt.Println(v.Name)
	}
}
