package pokedex

import "context"

//PokemonService is interface for pokedex service
type PokemonService interface {
	Create(ctx context.Context, pokemon Pokemon) (Pokemon, error)
	Fetch(ctx context.Context, filter Filter) (PokemonFeed, error)
	UpdateByID(ctx context.Context, ID string, pokemon Pokemon) (Pokemon, error)
	Delete(ctx context.Context, ID string) error
}

// PokemonRepository is interface for pokedex repository
type PokemonRepository interface {
	Create(ctx context.Context, pokemon Pokemon) (Pokemon, error)
	Fetch(ctx context.Context, filter Filter) ([]Pokemon, string, error)
	GetByID(ctx context.Context, ID string) (Pokemon, error)
	UpdateByID(ctx context.Context, ID string, pokemon Pokemon) (Pokemon, error)
	Delete(ctx context.Context, ID string) error
}
