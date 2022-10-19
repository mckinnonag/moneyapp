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
		private := v1.Group("/private")
		private.Use(middleware.Authz())
		{
			private.GET("/transactions", s.GetTransactions())
			private.POST("/transactions", s.CreateTransaction())
			private.GET("/contact", s.GetContact())
			private.POST("/contact", s.SetContact())
			private.DELETE("/contact", s.DeleteContact())
			plaid := private.Group("/plaid")
			{
				plaid.POST("/create_link_token", s.CreateLinkToken())
				plaid.POST("/set_access_token", s.GetAccessToken())
				plaid.GET("transactions", s.GetPlaidTransactions())
			}
		}
	}

	return router
}
