package pokemon

import (
	"context"

	"github.com/nugrohoac/pokedex"
	"github.com/pkg/errors"
)

type service struct {
	pokemonRepo pokedex.PokemonRepository
}

//Create is used to create pokemon
func (s service) Create(ctx context.Context, pokemon pokedex.Pokemon) (pokedex.Pokemon, error) {
	pokemons, _, err := s.pokemonRepo.Fetch(ctx, pokedex.Filter{Name: pokemon.Name})
	if err != nil {
		return pokedex.Pokemon{}, errors.Wrap(err, "error fetch pokemon by name at service")
	}

	if len(pokemons) > 0 {
		return pokedex.Pokemon{}, pokedex.ErrDuplicated{Message: "Pokemon already exist"}
	}

	return s.pokemonRepo.Create(ctx, pokemon)
}

// Fetch is used to fetch pokemon that can be filter
func (s service) Fetch(ctx context.Context, filter pokedex.Filter) (pokedex.PokemonFeed, error) {
	result := pokedex.PokemonFeed{}
	data, cursor, err := s.pokemonRepo.Fetch(ctx, filter)
	if err != nil {
		return pokedex.PokemonFeed{}, errors.Wrap(err, "error fetch pokemon")
	}

	result.Data.Pokemons = data
	result.Cursor = cursor

	return result, nil
}

// UpdateByID is used to update pokemon by id
func (s service) UpdateByID(ctx context.Context, ID string, pokemon pokedex.Pokemon) (pokedex.Pokemon, error) {
	_, err := s.pokemonRepo.GetByID(ctx, ID)
	if err != nil {
		return pokedex.Pokemon{}, errors.Wrap(err, "error get by id from pokemon repository")
	}

	return s.pokemonRepo.UpdateByID(ctx, ID, pokemon)
}

// Delete is function to delete pokemon by id
func (s service) Delete(ctx context.Context, ID string) error {
	_, err := s.pokemonRepo.GetByID(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "error get by id from pokemon repository")
	}

	return s.pokemonRepo.Delete(ctx, ID)
}

// NewPokemonService is constructor to initialize pokemon service
func NewPokemonService(pokemonRepo pokedex.PokemonRepository) pokedex.PokemonService {
	return service{
		pokemonRepo: pokemonRepo,
	}
}
