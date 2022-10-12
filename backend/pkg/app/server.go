package app

import (
	"log"
	"moneyapp/pkg/api"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router             *gin.Engine
	transactionService api.TransactionService
	plaidService       api.PlaidService
}

func NewServer(router *gin.Engine, transactionService api.TransactionService, plaidService api.PlaidService) *Server {
	return &Server{
		router:             router,
		transactionService: transactionService,
		plaidService:       plaidService,
	}
}

func (s *Server) Run(APP_PORT string) error {
	// run function that initializes the routes
	r := s.Routes()

	// run the server through the router
	err := r.Run(":" + APP_PORT)

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
