package models

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DATABASE_URL  string
	DATABASE_PORT int
	DATABASE_USER string
	DATABASE_NAME string
	DATABASE_PW   string
	DATABASE_SSL  string
)

func ConnectDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error when loading environment variables from .env file %w", err)
	}

	DATABASE_URL = os.Getenv("POSTGRES_URL")
	if DATABASE_URL == "" {
		DATABASE_URL = "localhost"
	}
	DATABASE_PORT, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	DATABASE_USER = os.Getenv("POSTGRES_USER")
	DATABASE_PW = os.Getenv("POSTGRES_PASSWORD")
	DATABASE_NAME = os.Getenv("POSTGRES_NAME")
	DATABASE_SSL = os.Getenv("POSTGRES_SSL")
	if DATABASE_USER == "" || DATABASE_PW == "" || DATABASE_NAME == "" {
		log.Fatal("DATABASE_USER or DATABASE_PW is not set")
	}
	if DATABASE_SSL == "" {
		DATABASE_SSL = "disable"
	}

	// Initialize DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		DATABASE_URL, DATABASE_PORT, DATABASE_USER, DATABASE_PW, DATABASE_NAME, DATABASE_SSL)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func init() {
	db := ConnectDB()
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("unable to ping database")
	}

	// DB Migration
	var migrationDir = flag.String("migration.files", "./migrations", "Directory where the migration files are located ?")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	file := "000003_remove_users_table.up.sql://../db/migrations"
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", *migrationDir),
		file, driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("Database Migrated!")
}
