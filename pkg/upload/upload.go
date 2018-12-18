package upload

type response struct {
	Status string `json:"status"`
	Name   string `json:"name"`
	Code   int16  `json:"code"`
}

const (
	UploadedFilePath = "/tmp/upload"
)
