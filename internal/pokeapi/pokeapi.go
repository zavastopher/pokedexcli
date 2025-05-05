package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const POKEAPI_ROOT_URL = "https://pokeapi.co/api/v2/"

type Config struct {
	Next     string
	Previous string
}

type LocationResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func LocationsRequest(confp *Config, resp *LocationResponse) (next string, prev string, err error) {
	conf := *confp
	res, err := http.Get(conf.Next)
	if err != nil {
		err = fmt.Errorf("Unable to complete location request: %v", err)
		return conf.Next, conf.Previous, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		err = fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return conf.Next, conf.Previous, err
	}

	err = json.Unmarshal(body, resp)
	if err != nil {
		err = fmt.Errorf("Unable to unmarshal data")
		return conf.Next, conf.Previous, err
	}

	return resp.Next, conf.Next, nil
}

func LocationsRequestBack(confp *Config, resp *LocationResponse) (next string, prev string, err error) {
	conf := *confp
	res, err := http.Get(conf.Previous)
	if err != nil {
		err = fmt.Errorf("Unable to complete location request: %v", err)
		return conf.Next, conf.Previous, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		err = fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return conf.Next, conf.Previous, err
	}

	err = json.Unmarshal(body, resp)
	if err != nil {
		err = fmt.Errorf("Unable to unmarshal data")
		return conf.Next, conf.Previous, err
	}

	if resp.Previous == nil || *resp.Previous == "" {
		return conf.Previous, "", nil
	}
	return conf.Previous, *resp.Previous, nil
}
