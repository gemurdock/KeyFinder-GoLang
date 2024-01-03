package util

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, templPath, tmpl string, data interface{}) {
	layout := filepath.Join(templPath, "layout.html")
	tmpl = filepath.Join(templPath, tmpl+".html")
	t, err := template.ParseFiles(layout, tmpl)
	if err != nil {
		fmt.Println("Error parsing template files...")
		fmt.Printf("Layout: %s\n", layout)
		fmt.Printf("Template: %s\n", tmpl)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("Error executing template...")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
