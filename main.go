package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gemurdock/KeyFinder-GoLang/api/middleware"
	"github.com/gemurdock/KeyFinder-GoLang/api/route"
	"github.com/gemurdock/KeyFinder-GoLang/db"
	"github.com/gemurdock/KeyFinder-GoLang/db/migrations"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func main() {
	// Get enviroment variables
	port := os.Getenv("APP_PORT")

	// Setup router
	r := setupRouter()

	// Connect to database
	dbi := db.GetDatabaseInstance()
	err := dbi.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	runStartupTasks(dbi.GetConnection())

	fmt.Println("Server running on port " + port)
	http.ListenAndServe(":"+port, r)
}

func runStartupTasks(db *gorm.DB) {
	migrations.MigrateAllModels(db)
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
