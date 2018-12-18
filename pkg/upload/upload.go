package upload

type response struct {
	Status string `json:"status"`
	Name   string `json:"name"`
	Code   int16  `json:"code"`
	Url    string `json:"url"`
}

const (
	UploadedFilePath = "/tmp/upload/"
)
