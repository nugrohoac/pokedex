package mysql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/nugrohoac/pokedex"

	"github.com/DATA-DOG/go-sqlmock"
)

//func (r *repositoryTest) TestRepoPokemon_Insert() {
//	r.T().Run("success", func(t *testing.T) {
//		repo := NewPokemonRepository(r.DB)
//
//		pokemonPayload := pokedex.Pokemon{
//			Number: "123",
//			Name:   "ratata",
//			Types:  []string{"grass"},
//		}
//		pokemonResponse, err := repo.Create(context.Background(), pokemonPayload)
//
//		require.NoError(t, err)
//		require.NotEmpty(t, pokemonResponse.ID)
//		require.Equal(t, pokemonPayload.Name, pokemonResponse.Name)
//		require.Equal(t, pokemonPayload.Types, pokemonResponse.Types)
//		require.Equal(t, pokemonPayload.Number, pokemonResponse.Number)
//	})
//}
//
//func (r *repositoryTest) TestRepoPokemon_Fetch() {
//	r.T().Run("success", func(t *testing.T) {
//
//		pokemon1 := pokedex.Pokemon{
//			Number: "123",
//			Name:   "Zygarde",
//			Types:  []string{"Dragon", "Ground"},
//		}
//		r.seedPokemon(pokemon1)
//
//		pokemon2 := pokedex.Pokemon{
//			Number: "111",
//			Name:   "Zapdos",
//			Types:  []string{"Electric", "Flying"},
//		}
//		r.seedPokemon(pokemon2)
//
//		repo := NewPokemonRepository(r.DB)
//		pokemonResponse, cursor, err := repo.Fetch(context.Background(), pokedex.Filter{
//			Limit: 1,
//		})
//
//		require.NoError(t, err)
//		require.NotEmpty(t, cursor)
//		require.Equal(t, 1, len(pokemonResponse))
//		require.Equal(t, pokemon2.Name, pokemonResponse[0].Name)
//	})
//}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("error sql mock")
	}

	return db, mock
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type AnyString struct{}

// Match satisfies sqlmock.Argument interface
func (as AnyString) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

func TestMysqlPokemonInsert(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		query := "INSERT INTO pokemon"
		db, _mock := NewMock()

		pokemon1 := pokedex.Pokemon{
			ID:        "id",
			Number:    "123",
			Name:      "pikachu",
			Types:     []string{"thunder"},
			CreatedAt: time.Now(),
		}

		_mock.ExpectExec(query).
			WithArgs(AnyString{}, pokemon1.Number, pokemon1.Name, strings.Join(pokemon1.Types, ","), AnyTime{}).
			WillReturnResult(sqlmock.NewResult(1, 1)).
			WillReturnError(nil)

		repo := NewPokemonRepository(db)
		resp, err := repo.Create(context.Background(), pokemon1)
		require.NoError(t, err)
		fmt.Println(resp)
	})
}
