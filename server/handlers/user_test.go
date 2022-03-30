package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	u := LoginPayload{
		User:     "user1@user.com",
		Password: "pass1",
	}

	payload, err := json.Marshal(&u)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/public/login", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	Login(c)

	assert.Equal(t, 200, w.Code)
}

func TestRegister(t *testing.T) {
	u := LoginPayload{
		User:     "user1@example.com",
		Password: "pass1",
	}

	payload, err := json.Marshal(&u)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/public/register", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	Register(c)

	assert.Equal(t, 200, w.Code)
}

func TestRegisterBadJSON(t *testing.T) {
	user := "test"

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/api/public/register", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	Register(c)

	assert.Equal(t, 400, w.Code)
}

func TestRegisterTakenUser(t *testing.T) {
	u1 := LoginPayload{
		User:     "user3@example.com",
		Password: "pass3",
	}
	u2 := LoginPayload{
		User:     "user3@example.com",
		Password: "pass3",
	}

	payload1, err := json.Marshal(&u1)
	assert.NoError(t, err)
	payload2, err := json.Marshal(&u2)
	assert.NoError(t, err)

	req1, err := http.NewRequest("POST", "/api/public/register", bytes.NewBuffer(payload1))
	assert.NoError(t, err)
	req2, err := http.NewRequest("POST", "api/public/register", bytes.NewBuffer(payload2))
	assert.NoError(t, err)

	w1 := httptest.NewRecorder()
	c1, _ := gin.CreateTestContext(w1)
	c1.Request = req1
	w2 := httptest.NewRecorder()
	c2, _ := gin.CreateTestContext(w2)
	c2.Request = req2

	Register(c1)
	assert.Equal(t, 200, w1.Code)
	Register(c2)
	assert.Equal(t, 400, w2.Code)
}

func TestLoginInvalidCredentials(t *testing.T) {
	user := LoginPayload{
		User:     "jwt@email.com",
		Password: "invalid",
	}

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/api/public/login", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	Login(c)

	assert.Equal(t, 401, w.Code)
}
