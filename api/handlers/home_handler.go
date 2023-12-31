package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type HomeData struct {
	TodaysNumber string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("template/layout.html", "template/"+tmpl+".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	currentTime := time.Now()
	dayOfMonth := strconv.Itoa(currentTime.Day())

	data := HomeData{
		TodaysNumber: dayOfMonth,
	}

	renderTemplate(w, "home", data)
}
