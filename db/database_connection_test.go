package db_test

import (
	"fmt"
	"testing"

	"github.com/gemurdock/KeyFinder-GoLang/db"
	"github.com/gemurdock/KeyFinder-GoLang/test"
)

var dbi *db.DatabaseConnection

func setup(t *testing.T) func(t *testing.T) {
	p := test.PostgresContainer{}
	err := p.Create()
	if err != nil {
		t.Errorf("Failed to create postgres container: %v", err)
		t.FailNow()
	}
	fmt.Println("Postgres container started")

	dbi = db.GetDatabaseInstance()
	dbi.LoadConfig(p.GetConfig())
	err = dbi.ConnectToDatabase()
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		t.FailNow()
	}
	fmt.Println("Database connection established")

	return func(t *testing.T) {
		dbi.CloseConnection()
		err = p.Destroy()
		if err != nil {
			t.Errorf("Failed to destroy postgres container: %v", err)
			t.FailNow()
		}
		fmt.Println("Postgres container destroyed")
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
