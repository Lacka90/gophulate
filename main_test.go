package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Init()
	code := m.Run()
	os.Exit(code)
}

func TestBadMode(t *testing.T) {
	req, _ := http.NewRequest("POST", "/process", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
	if body := response.Body.String(); !strings.Contains(body, "Process mode not supported") {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
