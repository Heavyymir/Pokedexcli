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
