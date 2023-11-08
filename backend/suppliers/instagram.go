package suppliers

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"shortformunifier/config"
)

const apiVersion = "v13.0"
const igUserID = "17841445329198685"

func UploadInstagram(cfg *config.Config, w http.ResponseWriter, request *http.Request, file io.Reader, videoReq *config.UploadVideoRequest) error {
	return createMediaContainer(videoReq)
}

func createMediaContainer(videoReq *config.UploadVideoRequest) error {
	url := "https://graph.facebook.com/" + apiVersion + igUserID + "/media"
	reelContainer := ReelContainer{
		MediaType:   "REELS",
		VideoURL:    "https://www.example.com/video.mp4",
		Caption:     videoReq.Title,
		ShareToFeed: true,
	}

	bytesRepresentation, err := json.Marshal(reelContainer)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	containerID := result["id"].(string)

	return publishVideo(containerID)
}

func publishVideo(containerID string) error {
	url := "https://graph.facebook.com/" + apiVersion + igUserID + "/media_publish"
	payload := map[string]string{
		"creation_id":  containerID,
		"access_token": "accessToken",
	}

	bytesRepresentation, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

type ReelContainer struct {
	MediaType             string   `json:"media_type"`
	VideoURL              string   `json:"video_url"`
	Caption               string   `json:"caption"`
	ShareToFeed           bool     `json:"share_to_feed"`
	CollaboratorUsernames []string `json:"collaborators"`
	CoverURL              string   `json:"cover_url"`
	AudioName             string   `json:"audio_name"`
	UserTags              []string `json:"user_tags"`
	LocationID            string   `json:"location_id"`
	ThumbOffset           int      `json:"thumb_offset"`
}
