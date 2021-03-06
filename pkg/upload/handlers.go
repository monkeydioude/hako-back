package upload

import (
	"github.com/monkeydioude/moon"
)

type dataResponse map[string]interface{}

type response struct {
	Status string       `json:"status"`
	Data   dataResponse `json:"data"`
	Code   int16        `json:"code"`
}

func Image(r *moon.Request) ([]byte, int, error) {
	f, h, err := r.HTTPRequest.FormFile("file")
	if err != nil {
		return []byte(err.Error()), 500, nil
	}
	defer f.Close()

	mimeType := h.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg":
		return saveImage(f, h.Filename, mimeType)
	case "image/png":
		return saveImage(f, h.Filename, mimeType)
	}
	return []byte("no mimetype found"), 404, nil
}

func DeleteImage(r *moon.Request) ([]byte, int, error) {
	if r.Matches["id"] == "" || r.Matches["user_id"] == "" {
		return jsonResponse404()
	}
	return deleteImage(r.Matches["id"], r.Matches["user_id"])
}
