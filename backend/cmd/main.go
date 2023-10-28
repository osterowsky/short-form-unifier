package main

import (
	"log"
	"net/http"
	"shortformunifier/pkg"
)

func main() {
	// We start the server running
	r := pkg.SetUpServer()

	pkg.SetUpRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
