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
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

func UploadYoutube(cfg *config.Config, w http.ResponseWriter, request *http.Request, file io.Reader, videoReq *config.VideoRequest) error {
	config := setupYoutubeConfig(cfg)
	token, err := loadTokenFromRefreshToken(config, request.Context())
	if err != nil {
		fmt.Printf("Error loading token: %v\n", err)
		return err
	}

	// Handle token expiry
	if err != nil && strings.Contains(err.Error(), "Token expired") {
		newToken, err := refreshAccessToken(config, token.RefreshToken)
		if err != nil {
			fmt.Printf("Error refreshing token: %v\n", err)
			return err
		}
		token = newToken

		// Save the updated token
		tokenJSON, _ := json.Marshal(token)
		ioutil.WriteFile("token.json", tokenJSON, 0644)
	}

	client := config.Client(context.Background(), token)

	youtubeService, err := youtube.New(client)
	if err != nil {
		fmt.Printf("Error creating YouTube client: %v\n", err)
		return err
	}

	ytSettings := videoReq.YoutubeConfig

	YouTubeRequest := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       videoReq.Title,
			Description: ytSettings.Description,
			Tags:        ytSettings.Tags,
		},
		Status: &youtube.VideoStatus{
			PrivacyStatus: videoReq.PrivacyLevel,
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

func loadTokenFromRefreshToken(config *oauth2.Config, ctx context.Context) (*oauth2.Token, error) {
	// Load token from a file or database
	tokenJSON, err := ioutil.ReadFile("token.json")
	if err != nil {
		return nil, err
	}

	var token oauth2.Token
	err = json.Unmarshal(tokenJSON, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func refreshAccessToken(config *oauth2.Config, refreshToken string) (*oauth2.Token, error) {
	ctx := context.Background()
	token := &oauth2.Token{
		RefreshToken: refreshToken,
		Expiry:       time.Now(), // Forces refresh
	}

	newToken, err := config.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, err
	}

	return newToken, nil
}

func setupYoutubeConfig(cfg *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.YoutubeClientID,
		ClientSecret: cfg.YoutubeClientSecret,
		RedirectURL:  "http://localhost",
		Scopes:       []string{"https://www.googleapis.com/auth/youtube.upload"},
		Endpoint:     google.Endpoint,
	}
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
