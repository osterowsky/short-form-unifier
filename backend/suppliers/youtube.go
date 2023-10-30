package suppliers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"google.golang.org/api/youtube/v3"
)

func UploadYoutube(w http.ResponseWriter, request *http.Request, file io.Reader) error {
	// Get youtube client
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	YouTubeRequest := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       "Test",
			Description: "Test",
			Tags:        []string{"test", "test2"},
		},
		Status: &youtube.VideoStatus{
			PrivacyStatus: "private",
		},
	}

	call := youtubeService.Videos.Insert([]string{"snippet,status"}, YouTubeRequest)

	// We need to modify to youtube file

	response, err := call.Media(file).Do()
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
	return nil
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
