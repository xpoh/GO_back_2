package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	Version = "1.0.0"
	Build   = "0000"
	Commit  = "unseе"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/__heartbeat__", heartbeatHandler)
	http.HandleFunc("/__version__", versionHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func heartbeatHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	response, err := json.Marshal(map[string]string{
		"version": Version,
		"commit":  Commit,
		"build":   Build,
	})
	_, err = fmt.Fprint(w, response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
