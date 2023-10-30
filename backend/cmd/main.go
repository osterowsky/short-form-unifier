package main

import (
	"log"
	"net/http"
	"shortformunifier/pkg"

	"github.com/rs/cors"
)

func main() {
	// We start the server running
	r := pkg.SetUpServer()

	pkg.SetUpRoutes(r)

	c := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", c))
}
