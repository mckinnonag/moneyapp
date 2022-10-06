package main

import (
	"fmt"
	"log"
	"server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	APP_PORT      string
	DATABASE_URL  string
	DATABASE_PORT int
	DATABASE_USER string
	DATABASE_NAME string
	DATABASE_PW   string
	DATABASE_SSL  string
)

// Holds the production app config
type server struct {
	router *gin.Engine
	// email  EmailSender
}

func init() {

}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error when loading environment variables from .env file %w", err)
	}

	s := &server{}
	err = models.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	err = s.routes()
	if err != nil {
		log.Fatal(err)
	}
}
