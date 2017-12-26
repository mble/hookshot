package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// RootEndpoint defines the behaviour for GET /
func RootEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("nothing here"))
}

// PingEndpoint defines the behaviour for GET /ping
func PingEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong!"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	router.HandleFunc("/ping", PingEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":80", router))
}
