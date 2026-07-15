package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

type locationAreaResponse struct {
	Next		string `json:"next"`
	Previous	*string `json:"previous"`
	Results		[]struct	{
			Name	string `json:"name"`
			URL		string `json:"url"`
	} `json:"results"`
}

func getPage(pageURL string)([]byte, error) {
	res, err := http.Get(pageURL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}
	return body, nil
}

func getLocationAreas(pageURL string) (locationAreaResponse, error) {
	body, err := getPage(pageURL)
	if err != nil {
		return locationAreaResponse{}, err
	}

	var data locationAreaResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		return locationAreaResponse{}, err
	}

	return data, nil
}
