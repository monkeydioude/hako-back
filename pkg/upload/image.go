package upload

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/monkeydioude/moon"
	"github.com/monkeydioude/tools"
)

const (
	ImageDirectory         = "img/"
	TmpUserId              = "0"
	TmpImageViewingBaseUrl = "localhost:8880"
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
		return saveImage(f, h.Filename)
	case "image/png":
		return saveImage(f, h.Filename)
	}
	return []byte("no mimetype found"), 404, nil
}

func getFileName(userID, name string) string {
	var ext string
	p := strings.Split(name, ".")
	if len(p) >= 2 {
		ext = p[1]
	}
	return fmt.Sprintf("%s.%s", tools.MD5(fmt.Sprintf("%s%d", TmpUserId, time.Now().Unix())).String(), ext)
}

func saveImage(file multipart.File, name string) ([]byte, int, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return jsonResponseErr(err.Error(), 500)
	}

	os.Mkdir(fmt.Sprintf("%s%s%s", UploadedFilePath, ImageDirectory, TmpUserId), 0766)

	fileName := getFileName(TmpUserId, name)
	fileExtendedPath := fmt.Sprintf("%s%s/%s", ImageDirectory, TmpUserId, fileName)
	fileURL := fmt.Sprintf("%s/%s", TmpImageViewingBaseUrl, fileExtendedPath)

	err = ioutil.WriteFile(fmt.Sprintf("%s%s", UploadedFilePath, fileExtendedPath), data, 0666)
	if err != nil {
		return jsonResponseErr(err.Error(), 500)
	}

	return jsonResponseOk(fileName, fileURL)
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

func jsonResponseOk(name, url string) ([]byte, int, error) {
	return jsonResponse(response{
		Status: "ok",
		Code:   200,
		Name:   name,
		Url:    url,
	}), 200, nil
}

func jsonResponseErr(status string, code int) ([]byte, int, error) {
	return jsonResponse(response{
		Status: status,
		Code:   int16(code),
	}), code, nil
}
