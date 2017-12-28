package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/mble/hookshot/docker"
	"github.com/mble/hookshot/version"
)

// RootEndpoint defines the behaviour for GET /
func RootEndpoint(w http.ResponseWriter, req *http.Request) {
	versionData := version.NewVersionData()
	render.Render(w, req, versionData)
}

// ImagesEndpoint defines the behaviour for GET /images
func ImagesEndpoint(w http.ResponseWriter, req *http.Request) {
	imgs := docker.ListImages()
	list := []render.Renderer{}
	for _, img := range imgs {
		list = append(list, img)
	}
	render.RenderList(w, req, list)
}

// DeployEndpoint defines the behaviour for GET /deploy
func DeployEndpoint(w http.ResponseWriter, req *http.Request) {
	imageName := chi.URLParam(req, "imageName")
	container, err := docker.DeployImage(imageName)
	if err != nil {
		panic(err)
	} else {
		render.Render(w, req, container)
	}
}

// New makes a new chi router with the right stuff
func New() *chi.Mux {
	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(Authenticate)

	router.Get("/", RootEndpoint)
	router.Get("/images", ImagesEndpoint)
	router.Post("/deploy/{imageName}", DeployEndpoint)
	return router
}
