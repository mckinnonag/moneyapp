package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	common "server/common"
	"strconv"
	"strings"
	"testing"
)

func getLoginPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "user1")
	params.Add("password", "pass1")

	return params.Encode()
}

func getRegistrationPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "u1")
	params.Add("password", "p1")

	return params.Encode()
}

func TestShowRegistrationPageUnauthenticated(t *testing.T) {
	r := common.GetRouter(true)

	r.GET("/u/register", ShowRegistrationPage)

	req, _ := http.NewRequest("GET", "/u/register", nil)

	common.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Register</title>") > 0

		return statusOK && pageOK
	})
}

func TestRegisterUnauthenticatedUnavailableUsername(t *testing.T) {
	common.SaveLists()
	w := httptest.NewRecorder()

	r := common.GetRouter(true)

	r.POST("/u/register", Register)

	registrationPayload := getLoginPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/register", strings.NewReader(registrationPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
	common.RestoreLists()
}

func TestRegisterUnauthenticated(t *testing.T) {
	common.SaveLists()
	w := httptest.NewRecorder()

	r := common.GetRouter(true)

	r.POST("/u/register", register)

	registrationPayload := getRegistrationPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/register", strings.NewReader(registrationPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Successful registration &amp; Login</title>") < 0 {
		t.Fail()
	}
	common.RestoreLists()
}
