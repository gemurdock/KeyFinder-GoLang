package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateAccountRoutes(r chi.Router) {
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Error: Not implemented"))
	})
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Error: Not implemented"))
	})
	r.Post("/logout", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Error: Not implemented"))
	})
	r.Post("/change-password", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Error: Not implemented"))
	})
	r.Get("/session-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Error: Not implemented"))
	})
	r.Post("/token-refresh", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Error: Not implemented"))
	})
}

func CreateAdminRoutes(r chi.Router) {
	r.Post("/reset-password", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Error: Not implemented"))
	})
	r.Post("/delete-account", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Error: Not implemented"))
	})
}
