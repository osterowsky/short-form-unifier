package suppliers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"shortformunifier/config"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

func UploadYoutube(cfg *config.Config, w http.ResponseWriter, request *http.Request, file io.Reader) error {
	config := &oauth2.Config{
		ClientID:     cfg.YoutubeClientID,
		ClientSecret: cfg.YoutubeClientSecret,
		RedirectURL:  "http://localhost",
		Scopes:       []string{"https://www.googleapis.com/auth/youtube.upload"},
		Endpoint:     google.Endpoint,
	}

	// Start the OAuth2 flow
	authURL := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v\n", authURL)

	var code string
	fmt.Print("Enter authorization code: ")
	fmt.Scan(&code)

	// Exchange the authorization code for a token
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		fmt.Printf("Error exchanging token: %v\n", err)
		return err
	}

	// Save the token to a file or database for future use
	tokenJSON, _ := json.Marshal(token)
	ioutil.WriteFile("token.json", tokenJSON, 0644)

	client := config.Client(context.Background(), token)

	youtubeService, err := youtube.New(client)
	if err != nil {
		fmt.Printf("Error creating YouTube client: %v\n", err)
		return err
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
