package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemurdock/KeyFinder-GoLang/db"
	"github.com/gemurdock/KeyFinder-GoLang/test"
)

var dbi *db.DatabaseConnection

func setup(t *testing.T) func(t *testing.T) { // todo: single setup/teardown for all tests
	conn, dbTeardown := test.SetupDBConn(t)
	dbi = conn
	return func(t *testing.T) {
		dbTeardown(t)
	}
}

func TestHomeHandler(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping this test in short mode")
	}

	teardown := setup(t)
	defer teardown(t)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	HomeHandler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}
