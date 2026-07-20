package pokeapi

import (
	"encoding/json"	
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := BaseURL + "pokemon/" + name

	body, err := c.get(url)
	if err != nil {
		return Pokemon{}, err
	}

	var data Pokemon

	err = json.Unmarshal(body, &data)
	if err != nil {
		return Pokemon{}, err
	}

	return data, nil
}
