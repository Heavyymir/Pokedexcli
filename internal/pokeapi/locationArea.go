package pokeapi

import (
	"encoding/json"
)

// Define the Areas for the API
type LocationAreaResponse struct {
	Next		string `json:"next"`
	Previous	*string `json:"previous"`
	Results		[]struct	{
			Name	string `json:"name"`
			URL		string `json:"url"`
	} `json:"results"`
}

// Define encounters for an Area
type LocationArea struct {
	Name		string `json:"name"`
	Encounters	[]Pokemon `json:"pokemon_encounters"`
}

// Get the locations for the user.
func (c *Client) GetLocationAreas(pageURL string) (LocationAreaResponse, error) {
	body, err := c.get(pageURL)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	var data LocationAreaResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return data, nil
}


//Get the encounters in a location for the user.
func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := BaseURL + "/" + name

	body, err := c.get(url)
	if err != nil {
		return LocationArea{}, err
	}

	var data LocationArea

	err = json.Unmarshal(body, &data)
	if err != nil {
		return LocationArea{}, err
	}

	return data, nil
}
