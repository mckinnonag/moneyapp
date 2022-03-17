package main

import (
	api "server/api"
	"server/common"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	common.Init()
}

func main() {
	r := gin.Default()

	// Load templates
	r.LoadHTMLGlob(common.TEMPLATE_PATH)

	corsConfig := cors.Default()
	// corsConfig.AllowOrigins = []string{"https://example.com"}
	// To be able to send tokens to the server.
	// corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	// corsConfig.AddAllowMethods("OPTIONS")
	r.Use(corsConfig)

	api.InitRoutes(r)

	err := r.Run(":" + common.APP_PORT)
	if err != nil {
		panic("unable to start server")
	}
}
