package pkg

import (
	"fmt"
	"net/http"
)

// We first just will test for Youtube and focus on that
// https://developers.google.com/youtube/v3/docs/videos/insert
// POST https://www.googleapis.com/upload/youtube/v3/videos
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("video")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Println(file)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully"))
}
