package db_test

import (
	"fmt"
	"testing"

	"github.com/gemurdock/KeyFinder-GoLang/db"
	"github.com/gemurdock/KeyFinder-GoLang/test"
)

var dbi *db.DatabaseConnection

func setup(t *testing.T) func(t *testing.T) {
	conn, dbTeardown := test.SetupDBConn(t)
	dbi = conn
	return func(t *testing.T) {
		dbTeardown(t)
	}
}

func Test_Database_Connection_With_Postgres(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping this test in short mode")
	}

	teardown := setup(t)
	defer teardown(t)

	fmt.Println("Pinging database")
	success, err := dbi.Ping()
	fmt.Println("Database pinged")

	if !success || err != nil {
		t.Errorf("Failed to ping database: %v", err)
		t.FailNow()
	}
}
