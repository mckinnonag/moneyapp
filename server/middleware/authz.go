package middleware

import (
	"fmt"
	"server/auth"

	"github.com/gin-gonic/gin"
)

// Authz validates token via Firebase and authorizes users
func Authz() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(403, "No Authorization header provided")
			c.Abort()
			return
		}

		bearerToken := c.Request.Header.Get("Authorization")
		_, err := auth.VerifyIDToken(c, bearerToken)
		if err != nil {
			if err.Error() == "illegal base64 data at input byte 6; see https://firebase.google.com/docs/auth/admin/verify-id-tokens for details on how to retrieve a valid ID token" {
				c.JSON(401, "Invalid authorization header.")
				c.Abort()
				return
			} else {
				c.JSON(400, "")
				fmt.Println(err.Error())
				c.Abort()
				return
			}
		}

		// c.Set("email", token.)

		c.Next()
	}
}
