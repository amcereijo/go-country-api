package app_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	app "github.com/amcereijo/go-country-api/app"
)

var a app.App

func TestMain(m *testing.M) {
	a.Intialize("../.env_test")

	code := m.Run()
	os.Exit(code)
}

func TestCall(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)

	fmt.Printf("Response %v\n", response)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body == "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
