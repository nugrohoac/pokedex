package http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/nugrohoac/pokedex"
	"github.com/pkg/errors"
)

type pokemonDelivery struct {
	pokemonService pokedex.PokemonService
	_validator     validator.Validate
	timeOut        time.Duration
}

// Create is used to store pokemon
// @Summary Store a new pokemon
// @Description Store a new pokemon
// @Tags root
// @Accept json
// @Produce json
// @Param pokemon body pokedex.Pokemon true "create pokemon"
// @Success 201 {object} pokedex.Pokemon "Success create a new pokemon"
// @Failure 409 {object} pokedex.ErrDuplicated "Pokemon with current name has been existing"
// @Router /pokemon [post]
func (p pokemonDelivery) Create(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), p.timeOut)
	defer cancel()

	var pokemon pokedex.Pokemon
	if err := c.Bind(&pokemon); err != nil {
		return errors.Wrap(pokedex.ErrBindStruct{Message: err.Error()}, "error bind struct pokemon")
	}

	if err := p._validator.Struct(pokemon); err != nil {
		return errors.Wrap(pokedex.ErrValidateStruct{Message: err.Error()}, "error validate struct pokemon")
	}

	storedPokemon, err := p.pokemonService.Create(ctx, pokemon)
	if err != nil {
		return errors.Wrap(err, "error store pokemon from pokemon service")
	}

	return c.JSON(http.StatusCreated, storedPokemon)
}

// Fetch is used to fetch pokemon by filter
// @Summary Fetch pokemon which filter by optional parameter
// @Description Fetch pokemon which filter by optional parameter
// @Tags root
// @Accept */*
// @Produce json
// @Param limit query integer false "limit items, default is 20"
// @Param cursor query string false "cursor is used to pagination"
// @Param name query string false "filter by name of pokemon"
// @Success 200 {object} pokedex.PokemonFeed
// @Router /pokemon [get]
func (p pokemonDelivery) Fetch(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), p.timeOut)
	defer cancel()

	filter := pokedex.Filter{
		Limit:  20,
		Cursor: c.QueryParam("cursor"),
		Name:   c.QueryParam("name"),
	}

	if limitStr := c.QueryParam("limit"); limitStr != "" {
		limitNumber, err := strconv.Atoi(limitStr)
		if err != nil {
			return errors.Wrap(pokedex.ErrInValid{Message: err.Error()}, "invalid value limit")
		}

		filter.Limit = int32(limitNumber)
	}

	pokemonFeed, err := p.pokemonService.Fetch(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "error fetch pokemon from pokemon service")
	}

	return c.JSON(http.StatusOK, pokemonFeed)
}

// UpdateByID is used to update existing pokemon
// @Summary Update pokemon by id
// @Description Update pokemon by id
// @Tags root
// @Accept json
// @Produce json
// @Param id path string true "id-pokemon"
// @Param pokemon body pokedex.Pokemon true "update pokemon"
// @Success 200 {object} pokedex.Pokemon
// @Failure 404 {object} pokedex.ErrNotFound "item will be update is not found"
// @Router /pokemon/{id} [put]
func (p pokemonDelivery) UpdateByID(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), p.timeOut)
	defer cancel()

	ID := c.Param("id")
	var pokemon pokedex.Pokemon
	if err := c.Bind(&pokemon); err != nil {
		return errors.Wrap(pokedex.ErrBindStruct{Message: err.Error()}, "error bind struct pokemon")
	}

	if err := p._validator.Struct(pokemon); err != nil {
		return errors.Wrap(pokedex.ErrValidateStruct{Message: err.Error()}, "error validate struct pokemon")
	}

	updatedPokemon, err := p.pokemonService.UpdateByID(ctx, ID, pokemon)
	if err != nil {
		return errors.Wrap(err, "error update pokemen from pokemon service")
	}

	return c.JSON(http.StatusOK, updatedPokemon)
}

// Delete pokemon by id
// @Summary Delete pokemon by id
// @Description Delete pokemon by id
// @Tags root
// @Accept json
// @Produce json
// @Param id path string true "id-pokemon"
// @Success 204
// @Failure 404 {object} pokedex.ErrNotFound "item will be delete is not found"
// @Router /pokemon/{id} [delete]
func (p pokemonDelivery) Delete(c echo.Context) error {
	ID := c.Param("id")

	ctx, cancel := context.WithTimeout(c.Request().Context(), p.timeOut)
	defer cancel()

	if err := p.pokemonService.Delete(ctx, ID); err != nil {
		return errors.Wrap(err, "error delete from pokemon service")
	}

	return c.NoContent(http.StatusNoContent)
}

// NewPokemonDelivery is used to initialized pokemon delivery
func NewPokemonDelivery(e *echo.Echo, pokemonService pokedex.PokemonService, timeOut time.Duration, _validator validator.Validate) {
	d := pokemonDelivery{
		pokemonService: pokemonService,
		timeOut:        timeOut,
		_validator:     _validator,
	}

	e.POST("/pokemon", d.Create)
	e.GET("/pokemon", d.Fetch)
	e.PUT("/pokemon/:id", d.UpdateByID)
	e.DELETE("/pokemon/:id", d.Delete)
}
