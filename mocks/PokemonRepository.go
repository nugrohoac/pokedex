// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	pokedex "github.com/nugrohoac/pokedex"
	mock "github.com/stretchr/testify/mock"
)

// PokemonRepository is an autogenerated mock type for the PokemonRepository type
type PokemonRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, pokemon
func (_m *PokemonRepository) Create(ctx context.Context, pokemon pokedex.Pokemon) (pokedex.Pokemon, error) {
	ret := _m.Called(ctx, pokemon)

	var r0 pokedex.Pokemon
	if rf, ok := ret.Get(0).(func(context.Context, pokedex.Pokemon) pokedex.Pokemon); ok {
		r0 = rf(ctx, pokemon)
	} else {
		r0 = ret.Get(0).(pokedex.Pokemon)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pokedex.Pokemon) error); ok {
		r1 = rf(ctx, pokemon)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, ID
func (_m *PokemonRepository) Delete(ctx context.Context, ID string) error {
	ret := _m.Called(ctx, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, filter
func (_m *PokemonRepository) Fetch(ctx context.Context, filter pokedex.Filter) ([]pokedex.Pokemon, string, error) {
	ret := _m.Called(ctx, filter)

	var r0 []pokedex.Pokemon
	if rf, ok := ret.Get(0).(func(context.Context, pokedex.Filter) []pokedex.Pokemon); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pokedex.Pokemon)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, pokedex.Filter) string); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, pokedex.Filter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, ID
func (_m *PokemonRepository) GetByID(ctx context.Context, ID string) (pokedex.Pokemon, error) {
	ret := _m.Called(ctx, ID)

	var r0 pokedex.Pokemon
	if rf, ok := ret.Get(0).(func(context.Context, string) pokedex.Pokemon); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(pokedex.Pokemon)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: ctx, ID, pokemon
func (_m *PokemonRepository) UpdateByID(ctx context.Context, ID string, pokemon pokedex.Pokemon) (pokedex.Pokemon, error) {
	ret := _m.Called(ctx, ID, pokemon)

	var r0 pokedex.Pokemon
	if rf, ok := ret.Get(0).(func(context.Context, string, pokedex.Pokemon) pokedex.Pokemon); ok {
		r0 = rf(ctx, ID, pokemon)
	} else {
		r0 = ret.Get(0).(pokedex.Pokemon)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, pokedex.Pokemon) error); ok {
		r1 = rf(ctx, ID, pokemon)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
