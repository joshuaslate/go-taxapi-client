package taxapi

import (
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	rootURI = "https://taxapi.io/api/"
	version = "1"
)

// Client is the client that will retrieve tax data from TaxAPI.io
type Client struct {
	cache         *cache.Cache
	httpClient    *http.Client
	nextRequestAt time.Time
}

// NewClient builds a client with given options to make requests to TaxAPI.io
func NewClient(autoCache bool) *Client {
	client := &Client{httpClient: http.DefaultClient, nextRequestAt: time.Now()}

	if autoCache {
		client.cache = cache.New(time.Hour*24, time.Minute)
	}

	return client
}
