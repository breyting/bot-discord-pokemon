package pokeapi

import (
	"net/http"
	"time"

	"github.com/breyting/pokedex-discord/pokedexcli/pokecache"
)

// Client -
type Client struct {
	Cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		Cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
