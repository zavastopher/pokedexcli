package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const POKEAPI_ROOT_URL = "https://pokeapi.co/api/v2/"

type config struct {
	next     string
	previous string
}

type LocationResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func locationsRequest(conf config) (next string, prev string, resp LocationResponse, err error) {
	res, err := http.Get(conf.next)
	if err != nil {
		fmt.Errorf("Unable to complete location request: %v", err)
		return "", "", LocationResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return "", "", LocationResponse{}, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Errorf("Unable to unmarshal data")
		return "", "", LocationResponse{}, err
	}
	return resp.Next, conf.next, resp, nil
}
