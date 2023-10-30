package suppliers

import (
	"context"
	"mime/multipart"
	"net/http"

	"google.golang.org/api/youtube/v3"
)

func UploadYoutube(w http.ResponseWriter, request *http.Request, file *multipart.File) error {

	// Get youtube client
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx)

	return err
}
