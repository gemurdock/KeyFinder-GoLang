package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gemurdock/KeyFinder-GoLang/config"
	"github.com/gemurdock/KeyFinder-GoLang/util"
)

type HomeData struct {
	TodaysNumber string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	currentTime := time.Now()
	dayOfMonth := strconv.Itoa(currentTime.Day())

	data := HomeData{
		TodaysNumber: dayOfMonth,
	}

	appDir, err := config.GetAppWorkingDir() // todo: move to util; before caused error due to recursion in imports
	if err != nil {
		fmt.Println("Error getting app working directory...")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmplPath := filepath.Join(appDir, "template")
	util.RenderTemplate(w, tmplPath, "home", data)
}
