package main

import (
	"fmt"
	"log"
	"net/http"
)

const version string = "1.0"

func root(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "sampleapp version %s", version)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "%s\n", version)
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/health", health)
	http.HandleFunc("/version", getVersion)
	log.Println("starting server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
