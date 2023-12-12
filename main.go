package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		fmt.Println("The current time is: ", currentTime, " - Hello World sent to the browser")
		w.Write([]byte("Hello World"))
	})
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
