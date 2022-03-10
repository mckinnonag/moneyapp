package common

import (
	"net/http"
	"net/http/httptest"
	"os"
	models "server/models"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []models.Article

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing
func GetRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// Helper function to process a request and test its response
func TestHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// This function is used to store the main lists into the temporary one for testing
func SaveLists() {
	// tmpArticleList = articleList
	// tmpUserList = userList
}

// This function is used to restore the main lists from the temporary one
func RestoreLists() {
	// articleList = tmpArticleList
	// userList = tmpUserList
}
