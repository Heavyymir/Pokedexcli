package pokeapi

import (
	"io"
	"net/http"
	"fmt"
)

func (c *Client) get(url string) ([]byte, error) {
	// Check to see if the cache has the data. If the data is found, return it.
	data, ok := c.cache.Get(url)
	if ok {
		return data, nil
	}

	// If not, create a new HTTP request. If this fails, return the error.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Send the newly created request.
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Check res statuscode. If the response code is above 299 (Failure) return an error.
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with statuscode: %d", res.StatusCode)
	}
	//Read the HTTP Body. If empty, return an error.
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	//Add the read request body data to the cache along with the URL.
	c.cache.Add(url, data)

	//Return the read request body.
	return data, nil
}
