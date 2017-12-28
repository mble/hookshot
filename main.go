package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mble/hookshot/router"
	"github.com/mble/hookshot/version"
)

func main() {
	router := router.New()
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9001"
	}
	fmt.Printf("Starting hookshot version: %s build: %s listening on: %s\n", version.VERSION, version.GITCOMMIT, port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
