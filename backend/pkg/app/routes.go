package app

import (
	"moneyapp/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func (s *Server) Routes() *gin.Engine {
	router := s.router

	// group all routes under /v1/api
	v1 := router.Group("/v1/api")
	{
		public := v1.Group("/public")
		{
			public.GET("/status", s.ApiStatus())
		}
		private := v1.Group("/private").Use(middleware.Authz())
		// private := v1.Group("/private")
		{
			private.POST("/transaction", s.CreateTransaction())
			private.GET("/get_link_token", s.CreateLinkToken())
			// private.POST("/linktoken", api.CreateLinkToken)
			// private.POST("/accesstoken", api.CreateAccessToken)
			// private.GET("/gettransactions", api.GetPlaidTransactions)
			// private.GET("/getsharedtransactions", api.GetTransactions)
			// private.GET("/accounts", api.GetAccounts)
			// private.POST("/removeaccount", api.RemoveAccount)
		}
	}

	return router
}
