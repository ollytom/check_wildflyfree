package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/icholy/digest"
)

// Read-only client of the Wildfly Managment HTTP API. It should be created with NewClient.
type Client struct {
	host string
	*http.Client
}

// MemUsage represents the JVM's heap memory usage in bytes.
type MemUsage struct {
	Used int
	Max  int
}

// NewClient creates a new read-only client of the Wildfly server at host authenticated using user and pass
// the Wildfly's authentication mechanism: HTTP digest authentication (yuck).
func NewClient(host, user, pass string) Client {
	return Client{
		host,
		&http.Client{
			Transport: &digest.Transport{
				Username: user,
				Password: pass,
			},
		},
	}
}

func (c *Client) StatMemory() (MemUsage, error) {
	u := url.URL{
		Scheme:   "http",
		Host:     c.host,
		Path:     "/management/core-service/platform-mbean/type/memory",
		RawQuery: "include-runtime=true&operation=attribute&name=heap-memory-usage",
	}
	resp, err := c.Get(u.String())
	if err != nil {
		return MemUsage{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return MemUsage{}, fmt.Errorf("%s: %s", resp.Status, resp.Body)
	}
	var mu MemUsage
	if err := json.NewDecoder(resp.Body).Decode(&mu); err != nil {
		return MemUsage{}, fmt.Errorf("decode response: %v", err)
	}
	return mu, nil
}
