package suppliers

import (
	"context"
	"mime/multipart"
	"net/http"
	"time"

	"google.golang.org/api/youtube/v3"
)

func UploadYoutube(w http.ResponseWriter, request *http.Request, file *multipart.File) error {

	// Get youtube client
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx)
	youtubeService.Videos.Insert([]string{"snippet,status"}, nil)

	return err
}

// Types for Youtube API
// ---------------------
type YoutubeVideoRequest struct {
	Snippet          *Snippet          `json:"snippet,omitempty"`
	Localizations    *Localizations    `json:"localizations,omitempty"`
	Status           *Status           `json:"status,omitempty"`
	RecordingDetails *RecordingDetails `json:"recordingDetails,omitempty"`
}

type Snippet struct {
	Title           string   `json:"title,omitempty"`
	Description     string   `json:"description,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	CategoryId      string   `json:"categoryId,omitempty"`
	DefaultLanguage string   `json:"defaultLanguage,omitempty"`
}

type Localizations struct {
	Key []*Key `json:"key,omitempty"`
}

type Key struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type Status struct {
	Embeddable              bool      `json:"embeddable,omitempty"`
	License                 string    `json:"license,omitempty"`
	PrivacyStatus           string    `json:"privacyStatus,omitempty"`
	PublicStatsViewable     bool      `json:"publicStatsViewable,omitempty"`
	PublishAt               time.Time `json:"publishAt,omitempty"`
	SelfDeclaredMadeForKids bool      `json:"selfDeclaredMadeForKids,omitempty"`
}

type RecordingDetails struct {
	RecordingDate time.Time `json:"recordingDate,omitempty"`
}
