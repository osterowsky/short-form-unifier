package suppliers

import (
	"fmt"
	"mime/multipart"
	"net/http"
)

func UploadYoutube(w http.ResponseWriter, request *http.Request, file *multipart.File) error {
	return fmt.Errorf("not implemented")
}
