package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Config holds the configuration parameters for the client.
type Config struct {
	Url  string // URL where the API is hosted
	Port string // Port number
}

// Client is an HTTP client that interacts with a remote API.
type Client struct {
	config Config       // Configuration for the client
	client *http.Client // HTTP client with timeout settings
}

// NewConfig creates a new Config object with the provided URL and port.
// url: The URL where the API is hosted.
// port: The port number on which the API is running.
func NewConfig(url string, port string) *Config {
	return &Config{url, port}
}

// NewClient creates a new Client object with the provided configuration.
// config: The configuration for the client, including URL and port.
func NewClient(config Config) *Client {
	return &Client{
		config: config,
		client: &http.Client{
			Timeout: time.Second * 10, // Timeout set to 10 seconds
		},
	}
}

// NilPokemon returns a default instance of Pokemon.
func NilPokemon() Pokemon {
	return NewPokemon("1", "a", "100")
}

// GetPokemonByName retrieves a Pokemon by its name from the API.
// Implements retry logic with exponential backoff.
// Returns the Pokemon and any error encountered.
func (c *Client) GetPokemonByName() (Pokemon, error) {
	url := c.config.Url
	var pokemon Pokemon

	// Retry logic for GetPokemonByName
	err := retry(3, 2*time.Second, func() (*http.Response, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		resp, err := c.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&pokemon)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})

	if err != nil {
		return NilPokemon(), err
	}

	return pokemon, nil
}

// GetPokemonList retrieves a list of Pokemon from the API.
// Implements retry logic with exponential backoff.
// Returns the list of Pokemon and any error encountered.
func (c *Client) GetPokemonList() (PokemonArray, error) {
	url := c.config.Url
	var pokemonArray PokemonArray

	// Retry logic for GetPokemonList
	err := retry(3, 2*time.Second, func() (*http.Response, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		resp, err := c.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&pokemonArray)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})

	if err != nil {
		return NewPokemonArray(), err
	}

	return pokemonArray, nil
}

// retry executes the provided function with retry logic and returns the response or error.
// attempts: The number of retry attempts.
// sleep: The duration to sleep between retry attempts.
// fn: The function to execute with retry logic.
func retry(attempts int, sleep time.Duration, fn func() (*http.Response, error)) error {
	for i := 0; i < attempts; i++ {
		_, err := fn()
		if err == nil {
			return nil
		}
		time.Sleep(sleep)
	}
	return fmt.Errorf("timed out waiting for response")
}
