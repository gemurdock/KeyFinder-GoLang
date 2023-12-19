package db_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/gemurdock/KeyFinder-GoLang/config"
	"github.com/gemurdock/KeyFinder-GoLang/db"
	"github.com/gemurdock/KeyFinder-GoLang/test"
)

func Test_Database_Connection_With_Postgres(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping this test in short mode")
	}

	ctx := context.Background()
	config := config.GetConfigInstance(false)
	config.LoadTestingValues()

	// Setup
	postgres := &test.PostgresContainer{}
	postgres.LoadConfig(config)
	err := postgres.Create(ctx)
	if err != nil {
		t.Errorf("Failed to create postgres container: %v", err)
		t.FailNow()
	}
	fmt.Println("Postgres container started")

	// Update config for container
	host, port := postgres.GetConnInfo()
	config.DBHost = host
	config.DBPort = port

	dbi := db.GetDatabaseInstance()
	dbi.LoadConfig(config)
	err = dbi.ConnectToDatabase()
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		t.FailNow()
	}
	defer dbi.CloseConnection()
	fmt.Println("Database connection established")

	fmt.Println("Pinging database")
	success, err := dbi.Ping()
	fmt.Println("Database pinged")

	if !success || err != nil {
		t.Errorf("Failed to ping database: %v", err)
		t.FailNow()
	}
}
