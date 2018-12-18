package upload

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"

	"github.com/monkeydioude/moon"
)

const (
	ImageDirectory = "/img"
)

func Image(r *moon.Request, c *moon.Configuration) ([]byte, int, error) {
	f, h, err := r.HTTPRequest.FormFile("file")
	if err != nil {
		return []byte(err.Error()), 500, nil
	}
	defer f.Close()

	mimeType := h.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg":
		return saveImage(f, h)
	case "image/png":
		return saveImage(f, h)
	}
	return []byte("no mimetype found"), 404, nil
}

func saveImage(file multipart.File, handle *multipart.FileHeader) ([]byte, int, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return jsonResponseErr(err.Error(), 500)

	}

	err = ioutil.WriteFile(fmt.Sprintf("%s%s/%s", UploadedFilePath, ImageDirectory, handle.Filename), data, 0666)
	if err != nil {
		return jsonResponseErr(err.Error(), 500)
	}
	return jsonResponseOk(handle.Filename)
}

func jsonResponse(data interface{}) []byte {
	res, err := json.Marshal(data)

	if err != nil {
		return []byte(`{
			"status": "ok",
			"code": 500
		}`)
	}

	return res
}

func jsonResponseOk(name string) ([]byte, int, error) {
	return jsonResponse(response{
		Status: "ok",
		Code:   200,
		Name:   name,
	}), 200, nil
}

func jsonResponseErr(status string, code int) ([]byte, int, error) {
	return jsonResponse(response{
		Status: status,
		Code:   int16(code),
	}), code, nil
}
