package app

import (
	"moneyapp/pkg/api"
	"moneyapp/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router             *gin.Engine
	l                  logger.Logger
	transactionService api.TransactionService
	plaidService       api.PlaidService
}

func NewServer(router *gin.Engine, l logger.Logger, transactionService api.TransactionService, plaidService api.PlaidService) *Server {
	return &Server{
		router:             router,
		l:                  l,
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
		return err
	}

	return nil
}
