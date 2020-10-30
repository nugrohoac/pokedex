package mysql

import (
	"database/sql"
	"path"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func migrateDB(db *sql.DB) (m *migrate.Migrate, err error) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, err
	}

	_, filename, _, _ := runtime.Caller(0)
	migrationPath := path.Join(path.Dir(filename), "migration")
	pathToMigrate := "file://" + migrationPath

	m, err = migrate.NewWithDatabaseInstance(pathToMigrate, "mysql", driver)
	if err != nil {
		return nil, err
	}
	err = m.Up()

	return
}
