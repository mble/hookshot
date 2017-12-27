package main

import (
	"log"
	"net/http"

	"github.com/mble/deployhook/router"
)

func main() {
	router := router.New()
	log.Fatal(http.ListenAndServe(":2015", router))
}
