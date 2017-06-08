package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/theRemix/Nineties-Namer/routes"
)

var r *mux.Router

func init() {
	r = routes.Router()
}

// RandomHandler should return a random name
// GET /random
func TestRandomHandler(t *testing.T) {

	req, _ := http.NewRequest("GET", "/random", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); len(body) == 0 {
		t.Errorf("Expected a non-zero length string. Got '%s'", body)
	}
}

// NamesHandler should return the same output for each input
// GET /names/{key}
func TestNameHandler(t *testing.T) {

	t.Run("GET /names/ should return 'fatboy kenan'", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "/names/", nil)
		response := executeRequest(req)

		checkResponseCode(t, http.StatusNotFound, response.Code)

	})

	t.Run("GET /names/something should return 'fatboy kenan'", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "/names/something", nil)
		response := executeRequest(req)

		checkResponseCode(t, http.StatusOK, response.Code)

		if body := response.Body.String(); body != "fatboy kenan" {
			t.Errorf("Expected response to be 'fatboy kenan'. Got '%s'", body)
		}

	})

	t.Run("GET /names/v1.0.3 should return 'wu tang girls' every time", func(t *testing.T) {

		req1, _ := http.NewRequest("GET", "/names/v1.0.3", nil)
		res1 := executeRequest(req1)

		req2, _ := http.NewRequest("GET", "/names/v1.0.3", nil)
		res2 := executeRequest(req2)

		body1 := res1.Body.String()
		body2 := res2.Body.String()

		if body1 != body2 {
			t.Errorf("Expected response to be 'wu tang girls' each time. Got '%s' and '%s'", body1, body2)
		}

	})
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	return rr
}
