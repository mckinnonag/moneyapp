package common

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DATABASE_URL  string
	DATABASE_PORT int
	DATABASE_USER string
	DATABASE_NAME string
	DATABASE_PW   string
	DATABASE_SSL  string
)

func Init() {
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
}
