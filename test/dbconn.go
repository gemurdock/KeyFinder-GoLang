package test

import (
	"fmt"
	"testing"

	"github.com/gemurdock/KeyFinder-GoLang/db"
)

func SetupDBConn(t *testing.T) (*db.DatabaseConnection, func(t *testing.T)) {
	p := PostgresContainer{}
	err := p.Create()
	if err != nil {
		t.Errorf("Failed to create postgres container: %v", err)
		t.FailNow()
	}
	fmt.Println("Postgres container started")

	dbi := db.GetDatabaseInstance()
	dbi.LoadConfig(p.GetConfig())
	err = dbi.ConnectToDatabase()
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
		t.FailNow()
	}

	return dbi, func(t *testing.T) {
		dbi.CloseConnection()
		err = p.Destroy()
		if err != nil {
			t.Errorf("Failed to destroy postgres container: %v", err)
			t.FailNow()
		}
		fmt.Println("Postgres container destroyed")
	}
}
