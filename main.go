package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.Handle("/", http.FileServer(http.Dir(".")))

	log.Printf("server started")
	log.Fatal(srv.ListenAndServe())
}
