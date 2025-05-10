package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const POKEAPI_ROOT_URL = "https://pokeapi.co/api/v2/"

type LocationResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func LocationsRequest(url string, resp *LocationResponse) (next string, prev string, results []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("Unable to complete location request: %v", err)
		return url, "", nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		err = fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return url, "", nil, err
	}

	err = json.Unmarshal(body, resp)
	if err != nil {
		err = fmt.Errorf("Unable to unmarshal data")
		return url, "", nil, err
	}

	return resp.Next, url, body, nil
}
