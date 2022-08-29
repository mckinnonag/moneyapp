package middleware

import (
	"server/auth"

	"github.com/gin-gonic/gin"
)

// Authz validates token and authorizes users
func Authz() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(403, "No Authorization header provided")
			c.Abort()
			return
		}

		// extractedToken := strings.Split(clientToken, "Bearer ")

		// if len(extractedToken) == 2 {
		// 	clientToken = strings.TrimSpace(extractedToken[1])
		// } else {
		// 	c.JSON(400, "Incorrect Format of Authorization Token")
		// 	c.Abort()
		// 	return
		// }

		bearerToken := c.Request.Header.Get("Authorization")
		token := auth.VerifyIDToken(c, bearerToken)
		if token == nil {
			c.JSON(400, "")
			c.Abort()
			return
		}

		// c.Set("email", token.)

		c.Next()

	}
}
