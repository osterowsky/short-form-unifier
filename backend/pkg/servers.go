package pkg

import (
	"github.com/gorilla/mux"
)

func SetUpServer() *mux.Router {
	r := mux.NewRouter()

	return r
}

func SetUpRoutes(r *mux.Router) {
	r.Handle("/", r)
	r.HandleFunc("/upload", UploadHandler).Methods("POST")
}
