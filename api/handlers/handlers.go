package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var (
	index, _  = template.ParseFiles("web/templates/index.html")
	err404, _ = template.ParseFiles("web/templates/err404.html")
	data      = PageData{
		Resultat: "",
		Title:    "Groupie-Tracker",
	}
)

// Handler of the main page
func HandleHome(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/favicon.ico" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Management of 404 error
	if r.URL.Path != "/" {
		err404.Execute(w, nil)
		log.Printf("Error 404: Page not found for path %s", r.URL.Path)
		return
	}

	err := index.Execute(w, data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	} else {
		log.Printf("Status OK : %v", http.StatusOK)
	}

}
