package main

import (
	"fmt"
	"net/http"

	"github.com/gemurdock/KeyFinder-GoLang/api/middleware"
	"github.com/gemurdock/KeyFinder-GoLang/api/route"
	"github.com/gemurdock/KeyFinder-GoLang/config"
	"github.com/gemurdock/KeyFinder-GoLang/db"
	"github.com/gemurdock/KeyFinder-GoLang/db/migrations"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func main() {
	config := config.GetConfigInstance(true)

	// Setup router
	r := setupRouter()

	// Connect to database
	dbi := db.GetDatabaseInstance()
	dbi.LoadConfig(config)
	err := dbi.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	runStartupTasks(dbi.GetConnection())

	fmt.Println("Server running on port " + config.AppPort)
	http.ListenAndServe(":"+config.AppPort, r)
}

func runStartupTasks(db *gorm.DB) {
	migrations.MigrateAllModels(db)
}

func setupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.ContentTypeSetter("application/json"))

	route.CreateAccountRoutes(r)
	route.CreateAdminRoutes(r)

	return r
}
