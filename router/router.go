package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mble/deployhook/docker"
	"github.com/mble/deployhook/version"
)

// RootEndpoint defines the behaviour for GET /
func RootEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(fmt.Sprintf("Version: %s Build: %s\n", version.VERSION, version.GITCOMMIT)))
}

// ImagesEndpoint defines the behaviour for GET /images
func ImagesEndpoint(w http.ResponseWriter, req *http.Request) {
	imgs := docker.ListImages()

	w.Write([]byte(imgs))
}

// DeployEndpoint defines the behaviour for GET /deploy
func DeployEndpoint(w http.ResponseWriter, req *http.Request) {
	imageName := chi.URLParam(req, "imageName")
	containerName, err := docker.DeployImage(imageName)
	if err != nil {
		panic(err)
	} else {
		msg := fmt.Sprintf("Deployed %s OK at %s\n", containerName, time.Now())
		fmt.Print(msg)
		w.Write([]byte(msg))
	}
}

// New makes a new chi router with the right stuff
func New() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/ping"))

	router.Get("/", RootEndpoint)
	router.Get("/images", ImagesEndpoint)
	router.Post("/deploy/{imageName}", DeployEndpoint)
	return router
}
