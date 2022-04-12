package models

import (
	"database/sql"
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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		"localhost", 5432, "postgres", "postgres", "data", "disable")

	var err error
	db, err := sql.Open("postgres", psqlInfo)
	assert.NoError(t, err)
	defer db.Close()

	err = db.Ping()
	assert.NoError(t, err)

	var migrationDir = flag.String("migration.files", "../db/migration", "Directory where the migration files are located ?")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", *migrationDir), // file://path/to/directory
		"000001_init_schema.up.sql://../db/migration", driver)
	assert.NoError(t, err)
	m.Up()

}
