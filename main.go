package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gemurdock/KeyFinder-GoLang/api/middleware"
	"github.com/gemurdock/KeyFinder-GoLang/api/route"
	"github.com/go-chi/chi/v5"
)

const (
	defaultPort = "3000"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.ContentTypeSetter("application/json"))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		fmt.Println("The current time is: ", currentTime, " - Hello World sent to the browser")
		w.Write([]byte("Hello World"))
	})

	route.CreateAccountRoutes(r)
	route.CreateAdminRoutes(r)

	fmt.Println("Server running on port " + defaultPort)
	http.ListenAndServe(":"+defaultPort, r)
}
