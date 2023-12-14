package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gemurdock/KeyFinder-GoLang/api/middleware"
	"github.com/gemurdock/KeyFinder-GoLang/api/route"
	"github.com/gemurdock/KeyFinder-GoLang/db/migrations"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Get enviroment variables
	port := os.Getenv("APP_PORT")

	// Setup router
	r := setupRouter()

	// Connect to database
	db, err := connectToDatabase()
	if err != nil {
		panic(err)
	}

	runStartupTasks(db)

	fmt.Println("Server running on port " + port)
	http.ListenAndServe(":"+port, r)
}

func runStartupTasks(db *gorm.DB) {
	migrations.MigrateAllModels(db)
}

func connectToDatabase() (*gorm.DB, error) {
	var maxRetries = 5
	var retryCount = 0

	// Retrieve environment variables
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")

	// Check if any of the required environment variables are missing
	if postgresUser == "" || postgresPassword == "" || postgresDB == "" {
		panic("Error: One or more PostgreSQL environment variables are missing.")
	}

	// Construct the connection string
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", postgresHost, postgresPort, postgresUser, postgresPassword, postgresDB)

	fmt.Printf("Connecting to %s:%s as %s to %s\n", postgresHost, postgresPort, postgresUser, postgresDB)

	var db *gorm.DB
	var err error

	for retryCount < maxRetries {
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
		if err != nil {
			retryCount++
			fmt.Printf("(%d of %d) Failed to connect to database. Retrying in 5 seconds. Error: %s\n", retryCount, maxRetries, err.Error())
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	if err != nil {
		return db, fmt.Errorf("failed to connect to database: %w", err)
	}
	fmt.Print("Successfully connected to database\n")

	return db, nil
}

func setupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.ContentTypeSetter("application/json"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		fmt.Println("The current time is: ", currentTime, " - Hello World sent to the browser")
		w.Write([]byte("Hello World"))
	})

	route.CreateAccountRoutes(r)
	route.CreateAdminRoutes(r)

	return r
}
