package pkg

import (
	"net/http"
	"shortformunifier/config"
	s "shortformunifier/suppliers"
)

// We first just will test for Youtube and focus on that
// https://developers.google.com/youtube/v3/docs/videos/insert
// POST https://www.googleapis.com/upload/youtube/v3/videos

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	cfg, err := config.NewConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("video")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// We are ready to send video for youtube
	err = s.UploadYoutube(cfg, w, r, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully"))
}
