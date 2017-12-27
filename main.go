package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mble/deployhook/docker"
	"github.com/mble/deployhook/version"
)

// RootEndpoint defines the behaviour for GET /
func RootEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(fmt.Sprintf("Version: %s Build: %s", version.VERSION, version.GITCOMMIT)))
}

// ImagesEndpoint defines the behaviour for GET /images
func ImagesEndpoint(w http.ResponseWriter, req *http.Request) {
	imgs := docker.ListImages()

	w.Write([]byte(imgs))
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/ping"))

	router.Get("/", RootEndpoint)
	router.Get("/images", ImagesEndpoint)
	log.Fatal(http.ListenAndServe(":2015", router))
}
