package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	messageConfigPath string = "/config/data"
	version           string = "1.0"
)

type response struct {
	Message string `json:"message"`
}

func root(w http.ResponseWriter, r *http.Request) {
	resp := &response{getMessage()}
	data, err := json.Marshal(resp)
	if err != nil {
		_, _ = fmt.Fprintf(w, "error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprint(w, string(data))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "%s\n", version)
}

func getMessage() string {
	defaultMsg := "Hello"
	if _, err := os.Stat(messageConfigPath); err == nil {
		data, err := ioutil.ReadFile(messageConfigPath)
		if err != nil {
			log.Printf("error reading message config: %s", err)
			return defaultMsg
		}
		return string(data)
	}
	return defaultMsg
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
