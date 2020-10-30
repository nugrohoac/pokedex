package mysql

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const dbNameTest = "pokemon_test"

type mysqlSuite struct {
	suite.Suite
	DB *sql.DB
	mg *migrate.Migrate
}

func (m *mysqlSuite) SetupSuite() {
	dsnDB := os.Getenv("MYSQL_DSN")
	if dsnDB == "" {
		dsnDB = "root:password123@tcp(127.0.0.1:3306)/" + dbNameTest + "?parseTime=1&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	}

	db, err := sql.Open("mysql", dsnDB)
	require.NoError(m.T(), err)
	require.NotNil(m.T(), db)

	m.mg, err = migrateDB(db)
	require.NoError(m.T(), err)

	m.DB = db
}

func (m *mysqlSuite) TearDownSuite() {
	require.NoError(m.T(), m.mg.Drop())
	require.NoError(m.T(), m.DB.Close())
}
