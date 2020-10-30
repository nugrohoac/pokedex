package pokedex

import (
	"time"
)

// Filter is used to filter fetch data
type Filter struct {
	Limit  int32
	Cursor string
	Name   string
}

// Pokemon is entity to describe pokedex object
type Pokemon struct {
	ID        string    `json:"id" example:"id-pokemon-string"`
	Number    string    `json:"number" validate:"required" example:"001"`
	Name      string    `json:"name" validate:"required" example:"pikachu"`
	Types     []string  `json:"types" validate:"required" example:"water, ice"`
	CreatedAt time.Time `json:"-"`
}

// PokemonFeed is entity for feed object response
type PokemonFeed struct {
	Cursor string      `json:"cursor"`
	Data   PokemonList `json:"data"`
}

// PokemonList .
type PokemonList struct {
	Pokemons []Pokemon `json:"pokemons"`
}
