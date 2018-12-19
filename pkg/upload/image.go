package upload

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/monkeydioude/hako-back/pkg/asset"
	"github.com/monkeydioude/hako-back/pkg/mongo"
	"github.com/monkeydioude/moon"
)

const (
	ImageDirectory         = "img/"
	TmpUserId              = "0"
	TmpImageViewingBaseUrl = "http://localhost:8880"
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

func saveImage(file multipart.File, name string) ([]byte, int, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return jsonResponseErr(err.Error(), 500)
	}

	os.Mkdir(fmt.Sprintf("%s%s%s", UploadedFilePath, ImageDirectory, TmpUserId), 0766)
	fileName, fileExtendedPath, fileURL := generateFileInfo(name)

	err = ioutil.WriteFile(fmt.Sprintf("%s%s", UploadedFilePath, fileExtendedPath), data, 0666)
	if err != nil {
		return jsonResponseErr(err.Error(), 500)
	}

	_, err = mongo.Database("test").Collection("test").InsertOne(&asset.Asset{
		URL:  fileURL,
		Type: "image",
	})
	if err != nil {
		return jsonResponseErr(err.Error(), 500)
	}

	return jsonResponseOk(fileName, fileURL)
}
