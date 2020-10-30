package mysql

import (
	"context"
	"database/sql"
	"encoding/base64"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"github.com/nugrohoac/pokedex"
)

type pokemonRepository struct {
	db *sql.DB
}

//Create is used to store pokemon
func (p pokemonRepository) Create(ctx context.Context, pokemon pokedex.Pokemon) (pokedex.Pokemon, error) {
	types := strings.Join(pokemon.Types, ",")

	pokemon.ID = uuid.NewV4().String()

	query, args, err := sq.Insert("pokemon").
		Columns(
			"id",
			"number",
			"name",
			"type",
			"created_at",
		).Values(
		pokemon.ID,
		pokemon.Number,
		pokemon.Name,
		types,
		time.Now(),
	).ToSql()

	if err != nil {
		return pokedex.Pokemon{}, errors.Wrap(err, "error building sql insert pokemon")
	}

	_, err = p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return pokedex.Pokemon{}, errors.Wrap(err, "error executing query insert pokemon")
	}

	return pokemon, nil
}

// Fetch is used to fetch pokemon by filter
func (p pokemonRepository) Fetch(ctx context.Context, filter pokedex.Filter) ([]pokedex.Pokemon, string, error) {
	pokemons := make([]pokedex.Pokemon, 0)
	fetchBuilder := sq.Select(
		"id",
		"number",
		"name",
		"type",
		"created_at",
	).From("pokemon").
		OrderBy("created_at DESC")

	if filter.Limit != 0 {
		fetchBuilder = fetchBuilder.Limit(uint64(filter.Limit))
	}

	if filter.Cursor != "" {
		cursor, err := decodeCursor(filter.Cursor)
		if err != nil {
			return pokemons, "", errors.Wrap(err, "error decode cursor")
		}

		fetchBuilder = fetchBuilder.Where("created_at < ?", cursor)
	}

	if filter.Name != "" {
		fetchBuilder = fetchBuilder.Where(sq.Eq{"name": filter.Name})
	}

	query, args, err := fetchBuilder.ToSql()
	if err != nil {
		return pokemons, "", errors.Wrap(err, "error convert builder to sql query")
	}

	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return pokemons, "", errors.Wrap(err, "error query row context fetch pokemon")
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			logrus.Error("error close connection : ", errClose)
		}
	}()

	for rows.Next() {
		pokemon := pokedex.Pokemon{}
		types := ""
		if err = rows.Scan(
			&pokemon.ID,
			&pokemon.Number,
			&pokemon.Name,
			&types,
			&pokemon.CreatedAt,
		); err != nil {
			logrus.Error("error scan pokemon : ", err)
			continue
		}

		pokemon.Types = strings.Split(types, ",")
		pokemons = append(pokemons, pokemon)
	}

	lenPokemons := len(pokemons)
	encodedcursor := ""
	if lenPokemons != 0 {
		cursor := pokemons[lenPokemons-1].CreatedAt
		encodedcursor, err = encodeCursor(cursor)

		if err != nil {
			return make([]pokedex.Pokemon, 0), encodedcursor, errors.Wrap(err, "error encode cursor")
		}
	}

	return pokemons, encodedcursor, nil
}

// GetByID is used to get pokemon by id
func (p pokemonRepository) GetByID(ctx context.Context, ID string) (pokedex.Pokemon, error) {
	query, args, err := sq.Select(
		"id",
		"number",
		"name",
		"type",
		"created_at",
	).From("pokemon").
		Where(sq.Eq{"id": ID}).
		ToSql()

	if err != nil {
		return pokedex.Pokemon{}, errors.Wrap(err, "error covert builder into query sql get by id pokemon")
	}

	row := p.db.QueryRowContext(ctx, query, args...)

	pokemon := pokedex.Pokemon{}
	types := ""
	err = row.Scan(
		&pokemon.ID,
		&pokemon.Number,
		&pokemon.Name,
		&types,
		&pokemon.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return pokedex.Pokemon{}, errors.Wrap(pokedex.ErrNotFound{Message: "pokemon is not found"}, "pokemon not found")
		}

		return pokedex.Pokemon{}, errors.Wrap(err, "error scan pokemon by id")
	}

	pokemon.Types = strings.Split(types, ",")
	return pokemon, nil
}

func (p pokemonRepository) UpdateByID(ctx context.Context, ID string, pokemon pokedex.Pokemon) (pokedex.Pokemon, error) {
	query, args, err := sq.Update("pokemon").
		Set("number", pokemon.Number).
		Set("name", pokemon.Name).
		Set("type", strings.Join(pokemon.Types, ",")).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": ID}).
		ToSql()

	if err != nil {
		return pokedex.Pokemon{}, errors.Wrap(err, "error building sql update pokemon")
	}

	_, err = p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return pokedex.Pokemon{}, errors.Wrap(err, "error executing query update pokemon")
	}

	return pokemon, nil
}

func (p pokemonRepository) Delete(ctx context.Context, ID string) error {
	query, args, err := sq.Delete("pokemon").
		Where(sq.Eq{"id": ID}).
		ToSql()

	if err != nil {
		return errors.Wrap(err, "error converting builder to sql delete pokemon")
	}

	_, err = p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "error executing delete pokemon")
	}

	return nil
}

// DecodeCursor takes a string and decode it using Base64 encoding to convert it
// into interface
func decodeCursor(cursor string) (interface{}, error) {
	byts, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", errors.Wrap(err, "error decode base64 cursor")
	}

	cursorTime, err := time.Parse(time.RFC3339, string(byts))
	if err != nil {
		return "", errors.Wrap(err, "error parse time into rfc3339")
	}

	return cursorTime, nil
}

// EncodeCursor takes a interface and encode its string representation
// using Base64 encoding by cursor type
func encodeCursor(cursor interface{}) (string, error) {
	byts, err := cursor.(time.Time).MarshalText()
	if err != nil {
		return "", errors.Wrap(err, "error casting cursor to time and marshall to text")
	}
	encodedcursor := base64.StdEncoding.EncodeToString(byts)

	return encodedcursor, nil
}

// NewPokemonRepository is used to initialize pokemon repository
func NewPokemonRepository(db *sql.DB) pokedex.PokemonRepository {
	return pokemonRepository{db: db}
}
