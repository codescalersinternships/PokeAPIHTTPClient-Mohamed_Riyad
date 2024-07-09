# PokeAPIHTTPClient

PokeAPIHTTPClient is a Go-based HTTP client designed to interact with the Pokémon API. The client provides functionalities to retrieve Pokémon data by name and fetch a list of Pokémon with retry logic for reliable communication.
## Table of Contents

    Installation
    Usage
    Configuration
    Endpoints
 
## Installation

To get started with the PokeAPIHTTPClient, you need to have Go installed on your system. You can download and install Go from the official website.
Clone the repository



    git clone https://github.com/codescalersinternships/PokeAPIHTTPClient-Mohamed_Riyad.git
    cd PokeAPIHTTPClient-Mohamed_Riyad

##Initialize the Go module



    go mod init github.com/codescalersinternships/PokeAPIHTTPClient-Mohamed_Riyad
    go mod tidy

## Usage

To use the PokeAPIHTTPClient, you need to configure it with the API URL and port. Below is an example of how to use the client to get Pokémon data by name and to fetch a list of Pokémon.

go

    package main
    
    import (
        "fmt"
        "log"
        "github.com/codescalersinternships/PokeAPIHTTPClient-Mohamed_Riyad/pkg"
    )
    
    func main() {
        config := pkg.NewConfig("https://pokeapi.co/api/v2/", "80")
        client := pkg.NewClient(*config)
    
        // Get Pokémon by name
        pokemon, err := client.GetPokemonByName("pikachu")
        if err != nil {
            log.Fatalf("Error getting Pokémon by name: %v", err)
        }
        fmt.Printf("Pokémon: %+v\n", pokemon)
    
        // Get Pokémon list
        pokemonList, err := client.GetPokemonList()
        if err != nil {
            log.Fatalf("Error getting Pokémon list: %v", err)
        }
        fmt.Printf("Pokémon List: %+v\n", pokemonList)
    }

## Configuration

The client requires a configuration object which holds the URL and port of the API.

go

    type Config struct {
        Url  string // URL where the API is hosted
        Port string // Port number
    }
    
    func NewConfig(url string, port string) *Config {
        return &Config{url, port}
    }

Endpoints

    GetPokemonByName: Fetches a Pokémon by its name.
    GetPokemonList: Retrieves a list of Pokémon.
