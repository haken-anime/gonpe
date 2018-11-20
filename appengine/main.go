package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ponpe/server/application"
	"github.com/ponpe/server/application/repository"
	"github.com/ponpe/server/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func init() {
	repos := repository.NewAllRepository()
	authApp := application.NewAuthenticationApp(
		repos,
	)

	http.Handle("/", handler.NewHandler(authApp))
}
