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
		{
			private.GET("/transaction", s.GetTransaction())
			private.POST("/transaction", s.CreateTransaction())
			private.POST("/create_link_token", s.CreateLinkToken())
			private.POST("/set_access_token", s.GetAccessToken())
		}
	}

	return router
}
