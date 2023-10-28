package pkg

import (
	"net/http"
)

// We first just will test for Youtube and focus on that
// https://developers.google.com/youtube/v3/docs/videos/insert
// POST https://www.googleapis.com/upload/youtube/v3/videos
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// test setup
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
