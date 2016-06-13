package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/huytd/api-togo"
)

func CreateTestServer() *httptest.Server {
	gin.SetMode(gin.ReleaseMode)
	server := main.Server{}
	app := &main.AppController{DB: server.InitDatabase("")}
	engine := server.InitEngine(app)
	testServer := httptest.NewServer(engine)
	return testServer
}

func TestGETEndpoints(t *testing.T) {
	var endpoints = []struct {
		in  string // input is an endpoint
		out int    // output is a status code
	}{
		{"", 200},                  // Request to index.html
		{"/api/", 200},             // Request to api index
		{"/api/map", 200},          // Get item list
		{"/api/map/test_key", 200}, // Get a key
		{"/api/map/test key", 400}, // Keys with space are not allowed
		{"/api/map/!@&", 400},      // Keys with special characters are not allowed
		{"/api/map/3575", 200},     // Keys with number are ok
		{"/api/map/api123", 200},   // Keys with alphanumerics are ok
		{"/api/map/123key", 200},   // Keys with alphanumerics and started with numbers is ok
	}

	server := CreateTestServer()
	defer server.Close()

	for _, endpoint := range endpoints {
		res, err := http.Get(server.URL + endpoint.in)
		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != endpoint.out {
			t.Error("GET request to " + endpoint.in + " Failed")
		}
	}
}
