package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

func main() {
	const filepathRoot = "."
	const port = ":8080"

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}

	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir(filepathRoot))
	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app/", fileserver))
	mux.Handle("/app/", fsHandler)
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /api/validate_chirp", validate_chirp)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("serving files on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
