package main

import (
	"handlers"
	"log"
	"net/http"
	"time"
)

func main() {

	fs := http.FileServer(http.Dir("./web/static/css"))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", fs))

	http.HandleFunc("/", handlers.Home)

	server := &http.Server{
		Addr:              ":8080",           // Address of the server (port is for example)
		ReadHeaderTimeout: 10 * time.Second,  // Time allowed to read headers
		WriteTimeout:      10 * time.Second,  // Max time to write response
		IdleTimeout:       120 * time.Second, // Max time between two requests
		MaxHeaderBytes:    1 << 20,           // 1 MB, maximum bytes server will read
	}
	log.Printf("Server starting on http://%s...\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
