package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mble/deployhook/version"
)

// RootEndpoint defines the behaviour for GET /
func RootEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(fmt.Sprintf("Version: %s Build: %s\n", version.VERSION, version.GITCOMMIT)))
}

// PingEndpoint defines the behaviour for GET /ping
func PingEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong!"))
}

func main() {
	router := chi.NewRouter()
	router.Get("/", RootEndpoint)
	router.Get("/ping", PingEndpoint)
	log.Fatal(http.ListenAndServe(":80", router))
}
