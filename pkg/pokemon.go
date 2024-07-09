package pkg

// Pokemon represents a single Pokémon with an ID, Name, and Score.
type Pokemon struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Score string `json:"score"`
}

// NewPokemon creates a new instance of Pokemon with the provided id, name, and score.
// id: The unique identifier for the Pokémon.
// name: The name of the Pokémon.
// score: The score associated with the Pokémon.
func NewPokemon(id string, name string, score string) Pokemon {
	return Pokemon{
		Id:    id,
		Name:  name,
		Score: score,
	}
}

// PokemonArray represents a collection of Pokemon.
type PokemonArray struct {
	PokemonArray []Pokemon `json:"pokemonList"`
}

// NewPokemonArray creates a new instance of PokemonArray with an empty list of Pokémon.
func NewPokemonArray() PokemonArray {
	return PokemonArray{}
}
