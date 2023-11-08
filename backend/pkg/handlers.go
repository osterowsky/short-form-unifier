package pkg

import (
	"encoding/json"
	"net/http"
	"shortformunifier/config"
	s "shortformunifier/suppliers"
)

// UploadHandler handles the upload of the video for all suppliers
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var videoReq config.UploadVideoRequest

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

	jsonData := r.FormValue("data")
	if err := json.Unmarshal([]byte(jsonData), &videoReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// We are ready to send video for youtube
	err = s.UploadYoutube(cfg, w, r, file, &videoReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// We also send video for tiktok
	err = s.UploadTiktok(cfg, w, r, file, &videoReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// And lastly for instagram
	err = s.UploadInstagram(cfg, w, r, file, &videoReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully"))
}
