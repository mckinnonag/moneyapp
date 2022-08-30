package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// type LoginPayload struct {
// 	User     string `json:"username"` //binding:"required"`
// 	Password string `json:"password"` //binding:"required"`
// }

// func Register(c *gin.Context) {
// 	var creds LoginPayload
// 	err := c.BindJSON(&creds)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(500, gin.H{
// 			"msg": "invalid json",
// 		})
// 	}

// 	if err := models.RegisterNewUser(creds.User, creds.Password); err == nil {
// 		j := auth.JwtConfig{
// 			SecretKey:  common.JWT_SECRET,
// 			Issuer:     common.JWT_ISSUER,
// 			Expiration: int64(common.JWT_EXPIRY),
// 		}

// 		token, err := j.GenerateToken(creds.User)
// 		if err != nil {
// 			log.Println(err)
// 			c.JSON(500, gin.H{
// 				"msg": "error generating token",
// 			})
// 		}

// 		type s struct {
// 			Token string `json:"token"`
// 		}
// 		obj := &s{
// 			Token: token,
// 		}
// 		data, _ := json.Marshal(obj)
// 		c.Data(200, "application/json", data)

// 	} else {
// 		c.JSON(400, err)
// 	}
// }

// func Login(c *gin.Context) {
// 	var creds LoginPayload
// 	err := c.BindJSON(&creds)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(500, gin.H{
// 			"msg": "invalid json",
// 		})
// 	}

// 	if models.IsUserValid(creds.User, creds.Password) {
// 		j := auth.JwtConfig{
// 			SecretKey:  common.JWT_SECRET,
// 			Issuer:     common.JWT_ISSUER,
// 			Expiration: int64(common.JWT_EXPIRY),
// 		}

// 		token, err := j.GenerateToken(creds.User)
// 		if err != nil {
// 			log.Println(err)
// 			c.JSON(500, gin.H{
// 				"msg": "error generating token",
// 			})
// 		}

// 		type s struct {
// 			Token string `json:"token"`
// 		}
// 		obj := &s{
// 			Token: token,
// 		}
// 		data, _ := json.Marshal(obj)
// 		c.Data(200, "application/json", data)
// 	} else {
// 		c.Data(401, "application/json", nil)
// 	}
// }

func Test(c *gin.Context) {
	fmt.Println("success!")
}
