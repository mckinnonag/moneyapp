package handlers

import (
	"encoding/json"
	common "server/common"
	models "server/models"

	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func generateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

func ShowRegistrationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	common.Render(c, gin.H{
		"title": "Register", "is_logged_in": c.MustGet("is_logged_in").(bool)}, "register.html")
}

func Register(c *gin.Context) {
	// Obtain the POSTed username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := models.RegisterNewUser(username, password); err == nil {
		// If the user is created, set the token in a cookie and log the user in
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		common.Render(c, gin.H{
			"title": "Successful registration & Login", "is_logged_in": c.MustGet("is_logged_in").(bool)}, "login-successful.html")

	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error(), "is_logged_in": c.MustGet("is_logged_in").(bool)})

	}
}

func ShowLoginPage(c *gin.Context) {
	common.Render(c, gin.H{
		"title": "Login", "is_logged_in": c.MustGet("is_logged_in").(bool),
	}, "login.html")
}

func PerformLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if models.IsUserValid(username, password) {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)

		common.Render(c, gin.H{
			"title": "Successful Login", "is_logged_in": c.MustGet("is_logged_in").(bool)}, "login-successful.html")

	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func ReactLogin(c *gin.Context) {
	// Test - send a dummy token
	data := []byte(`{"token":"test123"}`)
	c.Data(
		http.StatusOK,
		"application/json",
		data,
	)
}

// For testing
func ReactPerformLogin(c *gin.Context) {
	// username := c.PostForm("username")
	// password := c.PostForm("password")

	type LOGIN struct {
		USER     string `json:"username"` //binding:"required"`
		PASSWORD string `json:"password"` //binding:"required"`
	}

	var credentials LOGIN
	err := c.BindJSON(&credentials)
	if err != nil {
		panic(err)
	}

	type s struct {
		Token string `json:"token"`
	}
	obj := &s{
		Token: "test123",
	}
	data, _ := json.Marshal(obj)

	c.Data(http.StatusOK, "application/json", data)

	// if models.IsUserValid(credentials.USER, credentials.PASSWORD) {
	// 	token := generateSessionToken()
	// 	c.SetCookie("token", token, 3600, "", "", false, true)

	// 	common.Render(c, gin.H{
	// 		"title": "Successful Login", "is_logged_in": c.MustGet("is_logged_in").(bool)}, "login-successful.html")

	// } else {
	// 	c.HTML(http.StatusBadRequest, "login.html", gin.H{
	// 		"ErrorTitle":   "Login Failed",
	// 		"ErrorMessage": "Invalid credentials provided"})
	// }
}
