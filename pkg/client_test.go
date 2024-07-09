package pkg

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNewClient tests the creation of a new Client.
func TestNewClient(t *testing.T) {
	config := Config{Url: "http://localhost", Port: "8080"}
	client := NewClient(config)
	if client.config.Url != "http://localhost" || client.config.Port != "8080" {
		t.Errorf("Expected URL: %s and Port: %s, but got URL: %s and Port: %s", "http://localhost", "8080", client.config.Url, client.config.Port)
	}
}

// TestNewConfig tests the creation of a new Config.
func TestNewConfig(t *testing.T) {
	config := NewConfig("http://localhost", "8080")
	if config.Url != "http://localhost" || config.Port != "8080" {
		t.Errorf("Expected URL: %s and Port: %s, but got URL: %s and Port: %s", "http://localhost", "8080", config.Url, config.Port)
	}
}

// TestGetPokemonByName tests the GetPokemonByName method of the Client.
func TestGetPokemonByName(t *testing.T) {
	// Create a mock server.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokemon := NewPokemon("1", "Pikachu", "100")
		json.NewEncoder(w).Encode(pokemon)
	}))

	// Use the mock server URL in the client config.
	config := Config{Url: server.URL, Port: ""}
	client := NewClient(config)

	// Call the GetPokemonByName method.
	pokemon, err := client.GetPokemonByName()
	if err != nil {
		t.Errorf("error making request: %v", err)
	}

	// Check the response values.
	if pokemon.Id != "1" || pokemon.Name != "Pikachu" || pokemon.Score != "100" {
		t.Errorf("Unexpected pokemon: %v", pokemon)
	}
}

// TestGetPokemonList tests the GetPokemonList method of the Client.
func TestGetPokemonList(t *testing.T) {
	// Create a mock server.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokemonList := PokemonArray{
			PokemonArray: []Pokemon{
				NewPokemon("1", "Pikachu", "100"),
				NewPokemon("2", "Charmander", "200"),
			},
		}
		json.NewEncoder(w).Encode(pokemonList)
	}))

	// Use the mock server URL in the client config.
	config := Config{Url: server.URL, Port: ""}
	client := NewClient(config)

	// Call the GetPokemonList method.
	pokemonArray, err := client.GetPokemonList()
	if err != nil {
		t.Errorf("error making request: %v", err)
	}

	// Check the response values.
	if len(pokemonArray.PokemonArray) != 2 {
		t.Errorf("Unexpected number of pokemons: %v", len(pokemonArray.PokemonArray))
	}
	if pokemonArray.PokemonArray[0].Id != "1" || pokemonArray.PokemonArray[0].Name != "Pikachu" || pokemonArray.PokemonArray[0].Score != "100" {
		t.Errorf("Unexpected pokemon: %v", pokemonArray.PokemonArray[0])
	}
	if pokemonArray.PokemonArray[1].Id != "2" || pokemonArray.PokemonArray[1].Name != "Charmander" || pokemonArray.PokemonArray[1].Score != "200" {
		t.Errorf("Unexpected pokemon: %v", pokemonArray.PokemonArray[1])
	}
}
