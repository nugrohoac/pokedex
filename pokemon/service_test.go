package pokemon_test

import (
	"context"
	"errors"
	"testing"

	_errors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/nugrohoac/pokedex"
	"github.com/nugrohoac/pokedex/mocks"
	"github.com/nugrohoac/pokedex/pokemon"
	"github.com/nugrohoac/pokedex/testdata"
)

func TestService_Create(t *testing.T) {
	articuno := pokedex.Pokemon{
		ID:     "d3244c2f-7fea-4887-b7bb-1c7f274907d5",
		Number: "0123",
		Name:   "Articuno",
		Types:  []string{"Ice", "Flying"},
	}
	tests := map[string]struct {
		pokemonParam   pokedex.Pokemon
		createPokemon  testdata.FuncCaller
		expectedResult pokedex.Pokemon
		expectedError  error
	}{
		"success": {
			pokemonParam: articuno,
			createPokemon: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, articuno},
				Output:   []interface{}{articuno, nil},
			},
			expectedResult: articuno,
			expectedError:  nil,
		},
		"failed": {
			pokemonParam: articuno,
			createPokemon: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, articuno},
				Output:   []interface{}{pokedex.Pokemon{}, errors.New("failed store pokemon")},
			},
			expectedResult: pokedex.Pokemon{},
			expectedError:  errors.New("failed store pokemon"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			pokemonRepoMock := new(mocks.PokemonRepository)
			pokemonService := pokemon.NewPokemonService(pokemonRepoMock)

			if test.createPokemon.IsCalled {
				pokemonRepoMock.On("Create", test.createPokemon.Input...).
					Return(test.createPokemon.Output...).Once()
			}

			result, err := pokemonService.Create(context.Background(), test.pokemonParam)

			pokemonRepoMock.AssertExpectations(t)

			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, test.expectedError, _errors.Cause(err))
		})
	}
}

func TestService_Fetch(t *testing.T) {
	pokemons := []pokedex.Pokemon{
		{
			ID:     "d3244c2f-7fea-4887-b7bb-1c7f274907d5",
			Number: "0123",
			Name:   "Articuno",
			Types:  []string{"Ice", "Flying"},
		},
		{
			ID:     "39fb4f6b-5d0a-4514-b4a3-3af493954765",
			Number: "0124",
			Name:   "Zapdos",
			Types:  []string{"Electric", "Flying"},
		},
	}

	tests := map[string]struct {
		paramFilter    pokedex.Filter
		fetchPokemon   testdata.FuncCaller
		expectedResult pokedex.PokemonFeed
		expectedErr    error
	}{
		"success with limit": {
			paramFilter: pokedex.Filter{
				Limit: 2,
			},
			fetchPokemon: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, pokedex.Filter{Limit: 2}},
				Output:   []interface{}{pokemons, "next-cursor", nil},
			},
			expectedResult: pokedex.PokemonFeed{
				Cursor: "next-cursor",
				Data: pokedex.PokemonList{
					Pokemons: pokemons,
				},
			},
			expectedErr: nil,
		},
		"error fetch with limit": {
			paramFilter: pokedex.Filter{
				Limit: 2,
			},
			fetchPokemon: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, pokedex.Filter{Limit: 2}},
				Output:   []interface{}{[]pokedex.Pokemon{}, "", errors.New("error fetch pokemon")},
			},
			expectedResult: pokedex.PokemonFeed{},
			expectedErr:    errors.New("error fetch pokemon"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			pokemonRepoMock := new(mocks.PokemonRepository)
			pokemonService := pokemon.NewPokemonService(pokemonRepoMock)

			if test.fetchPokemon.IsCalled {
				pokemonRepoMock.On("Fetch", test.fetchPokemon.Input...).
					Return(test.fetchPokemon.Output...).Once()
			}

			result, err := pokemonService.Fetch(context.Background(), test.paramFilter)
			pokemonRepoMock.AssertExpectations(t)

			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, test.expectedErr, _errors.Cause(err))
		})
	}
}

func TestService_UpdateByID(t *testing.T) {
	articuno := pokedex.Pokemon{
		ID:     "d3244c2f-7fea-4887-b7bb-1c7f274907d5",
		Number: "0123",
		Name:   "Articuno",
		Types:  []string{"Ice", "Flying"},
	}
	articunoUpdate := articuno
	articunoUpdate.Name = "Articuno Update"
	articunoUpdate.Number = "0135"

	tests := map[string]struct {
		paramID        string
		paramPokemon   pokedex.Pokemon
		getPokemonByID testdata.FuncCaller
		updatePokemon  testdata.FuncCaller
		expectedResult pokedex.Pokemon
		expectedErr    error
	}{
		"success": {
			paramID:      articuno.ID,
			paramPokemon: articuno,
			getPokemonByID: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, articuno.ID},
				Output:   []interface{}{articuno, nil},
			},
			updatePokemon: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, articuno.ID, articuno},
				Output:   []interface{}{articunoUpdate, nil},
			},
			expectedResult: articunoUpdate,
			expectedErr:    nil,
		},
		"error get by id": {
			paramID:      articuno.ID,
			paramPokemon: articuno,
			getPokemonByID: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, articuno.ID},
				Output:   []interface{}{pokedex.Pokemon{}, errors.New("error get pokemon by id")},
			},
			updatePokemon:  testdata.FuncCaller{},
			expectedResult: pokedex.Pokemon{},
			expectedErr:    errors.New("error get pokemon by id"),
		},
		"failed": {
			paramID:      articuno.ID,
			paramPokemon: articuno,
			getPokemonByID: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, articuno.ID},
				Output:   []interface{}{articuno, nil},
			},
			updatePokemon: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, articuno.ID, articuno},
				Output:   []interface{}{pokedex.Pokemon{}, errors.New("failed update pokemon by id")},
			},
			expectedResult: pokedex.Pokemon{},
			expectedErr:    errors.New("failed update pokemon by id"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			pokemonRepoMock := new(mocks.PokemonRepository)
			pokemonService := pokemon.NewPokemonService(pokemonRepoMock)

			if test.getPokemonByID.IsCalled {
				pokemonRepoMock.On("GetByID", test.getPokemonByID.Input...).
					Return(test.getPokemonByID.Output...).Once()
			}

			if test.updatePokemon.IsCalled {
				pokemonRepoMock.On("UpdateByID", test.updatePokemon.Input...).
					Return(test.updatePokemon.Output...).Once()
			}

			result, err := pokemonService.UpdateByID(context.Background(), test.paramID, test.paramPokemon)
			pokemonRepoMock.AssertExpectations(t)

			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, test.expectedErr, _errors.Cause(err))
		})
	}
}

func TestService_Delete(t *testing.T) {
	articuno := pokedex.Pokemon{
		ID:     "d3244c2f-7fea-4887-b7bb-1c7f274907d5",
		Number: "0123",
		Name:   "Articuno",
		Types:  []string{"Ice", "Flying"},
	}

	tests := map[string]struct {
		paramID        string
		getPokemonByID testdata.FuncCaller
		deletePokemon  testdata.FuncCaller
		expectedErr    error
	}{
		"success": {
			paramID: articuno.ID,
			getPokemonByID: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, articuno.ID},
				Output:   []interface{}{articuno, nil},
			},
			deletePokemon: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{mock.Anything, articuno.ID},
				Output:   []interface{}{nil},
			},
			expectedErr: nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			pokemonRepoMock := new(mocks.PokemonRepository)
			pokemonService := pokemon.NewPokemonService(pokemonRepoMock)

			if test.getPokemonByID.IsCalled {
				pokemonRepoMock.On("GetByID", test.getPokemonByID.Input...).
					Return(test.getPokemonByID.Output...).Once()
			}

			if test.deletePokemon.IsCalled {
				pokemonRepoMock.On("Delete", test.deletePokemon.Input...).
					Return(test.deletePokemon.Output...).Once()
			}

			err := pokemonService.Delete(context.Background(), articuno.ID)
			pokemonRepoMock.AssertExpectations(t)

			assert.Equal(t, test.expectedErr, _errors.Cause(err))
		})
	}
}
