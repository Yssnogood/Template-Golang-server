package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Error string
}

const webTitle string = "Web-Template"

func RenderTemplate(w http.ResponseWriter, tmpl string, page Page) {
	t, err := template.ParseFiles("./web/templates/" + tmpl + ".html")

	if err != nil {
		ErrorPage(w, http.StatusBadRequest, Page{Title: webTitle, Error: "Error 400"})
		return
	}

	if err := t.Execute(w, page); err != nil {
		ErrorPage(w, http.StatusBadRequest, Page{Title: webTitle, Error: "Error 400"})
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/favicon.ico" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.URL.Path != "/" {
		ErrorPage(w, http.StatusNotFound, Page{Title: webTitle, Error: "Error 404"})
		return
	}

	switch r.Method {
	case "GET":
		RenderTemplate(w, "Index", Page{Title: webTitle})
		log.Printf("HTTP Response Code : %v", (http.StatusOK))
	default:
		ErrorPage(w, http.StatusMethodNotAllowed, Page{Title: webTitle, Error: "Error 405"})
	}
}

func ErrorPage(w http.ResponseWriter, errorCode int, page Page) {
	RenderTemplate(w, "Error", page)
	log.Printf("HTTP Response Code : %v", errorCode)
}
