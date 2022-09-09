package models

import (
	"flag"
	"fmt"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestDBConnection(t *testing.T) {
	db := ConnectDB()
	err := db.Ping()
	assert.NoError(t, err)

	var migrationDir = flag.String("migration.files", "../db/migrations", "Directory where the migration files are located ?")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	assert.NoError(t, err)
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", *migrationDir), // file://path/to/directory
		"000001_init_schema.up.sql://../db/migration", driver)
	assert.NoError(t, err)
	m.Up()
}
