package pokeapi

import (
	"net/http"
	"time"

	"github.com/Heavyymir/Pokedex.cli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		httpClient: http.Client{},
		cache:      pokecache.NewCache(5 * time.Second),
	}
}
