package suppliers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"shortformunifier/config"
)

// https://developers.tiktok.com/doc/content-posting-api-reference-direct-post/
func UploadTiktok(cfg *config.Config, w http.ResponseWriter, request *http.Request, file io.Reader, videoReq *config.VideoRequest) error {
	// Read the video content from the io.Reader
	videoContent, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	const chunkSize = 2 * 1024 * 1024 // 2MB in bytes
	// Calculate the total number of chunks
	totalChunks := len(videoContent) / chunkSize
	if len(videoContent)%chunkSize != 0 {
		totalChunks++
	}

	TiktokRequest := &TiktokVideoRequest{
		PostInfo: &PostInfo{
			Title: videoReq.Title,
		},
		SourceInfo: &SourceInfo{
			Source:          "FILE_UPLOAD",
			VideoSize:       int64(len(videoContent)),
			ChunkSize:       chunkSize,
			TotalChunkCount: int64(totalChunks),
		},
	}
	// We transform privacy level from our format to Tiktok's format
	TiktokRequest.PostInfo.PrivacyLevel = getTiktokPrivacyLevel(videoReq.PrivacyLevel)

	return uploadVideo(TiktokRequest)
}

func uploadVideo(tvr *TiktokVideoRequest) error {
	requestBody, err := json.Marshal(tvr)
	if err != nil {
		return err
	}

	// Create a new request
	req, err := http.NewRequest("POST", "https://open.tiktokapis.com/v2/post/publish/creator_info/query/", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer act.example12345Example12345Example")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// Make the API request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	fmt.Println(resp.Body)
	return nil

}

func getTiktokPrivacyLevel(privacyLevel string) string {
	switch privacyLevel {
	case "public":
		return "PUBLIC_TO_EVERYONE"
	case "private":
		return "SELF_ONLY"
	default:
		return "SELF_ONLY"
	}
}

type TiktokVideoRequest struct {
	PostInfo   *PostInfo   `json:"post_info"`
	SourceInfo *SourceInfo `json:"source_info"`
}

// PostInfo specifies detail about tiktok post
type PostInfo struct {
	PrivacyLevel          string `json:"privacy_level"`
	Title                 string `json:"title,omitempty"`
	DisableDuet           bool   `json:"disable_duet,omitempty"`
	DisableStich          bool   `json:"disable_stich,omitempty"`
	DisableComment        bool   `json:"disable_comment,omitempty"`
	VideoCoverTimeStampMs int32  `json:"video_cover_time_stamp_ms,omitempty"`
	BrandContentToggle    bool   `json:"brand_content_toggle,omitempty"`
	BrandOrganicToggle    bool   `json:"brand_organic_toggle,omitempty"`
}

// SourceInfo specifies details about video
type SourceInfo struct {
	Source          string `json:"source"`
	VideoURL        string `json:"video_url,omitempty"`
	VideoSize       int64  `json:"video_size"`
	ChunkSize       int64  `json:"chunk_size"`
	TotalChunkCount int64  `json:"total_chunk_count"`
}
