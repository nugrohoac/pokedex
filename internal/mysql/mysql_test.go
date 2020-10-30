package mysql

import (
	"strings"
	"testing"
	"time"

	"github.com/nugrohoac/pokedex"
	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type repositoryTest struct {
	mysqlSuite
}

func TestMYSQLRepository(t *testing.T) {
	if testing.Short() {
		t.Skip(`Skip comment repository test`)
	}

	suite.Run(t, new(repositoryTest))
}

func (r *repositoryTest) TearDownTest() {
	query := `SELECT CONCAT('TRUNCATE TABLE ', TABLE_NAME, ';') AS truncateCommand
		FROM information_schema.TABLES WHERE TABLE_SCHEMA = '` + dbNameTest + `';`

	rows, err := r.DB.Query(query)
	require.NoError(r.T(), err)
	require.NoError(r.T(), rows.Err())

	for rows.Next() {
		var query string
		err := rows.Scan(&query)
		require.NoError(r.T(), err)
		_, err = r.DB.Exec(query)
		require.NoError(r.T(), err)
	}

	rows.Close()
}

func (r *repositoryTest) seedPokemon(p pokedex.Pokemon) {
	if p.ID == "" {
		p.ID = uuid.NewV4().String()
	}

	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
	}

	_, err := r.DB.Exec(`INSERT INTO pokemon (id, number, name, type, created_at ) VALUES (?, ?, ?, ?, ?)`,
		p.ID,
		p.Number,
		p.Name,
		strings.Join(p.Types, ","),
		p.CreatedAt,
	)

	require.NoError(r.T(), err)
}
