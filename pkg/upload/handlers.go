package upload

import "github.com/monkeydioude/moon"

type dataResponse map[string]interface{}

type response struct {
	Status string       `json:"status"`
	Data   dataResponse `json:"data"`
	Code   int16        `json:"code"`
}

func Image(r *moon.Request, c *moon.Configuration) ([]byte, int, error) {
	f, h, err := r.HTTPRequest.FormFile("file")
	if err != nil {
		return []byte(err.Error()), 500, nil
	}
	defer f.Close()

	mimeType := h.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg":
		return saveImage(f, h.Filename)
	case "image/png":
		return saveImage(f, h.Filename)
	}
	return []byte("no mimetype found"), 404, nil
}
