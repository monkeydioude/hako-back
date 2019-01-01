package upload

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/monkeydioude/hako-back/pkg/asset"
	"github.com/monkeydioude/hako-back/pkg/mongo"
)

const (
	ImageDirectory = "img/"
	TmpUserId      = "0"
)

func saveImage(file multipart.File, name, mimeType string) ([]byte, int, error) {
	data, err := ioutil.ReadAll(file)
	UID := TmpUserId
	if err != nil {
		return jsonResponseErr(err, 500)
	}

	// Might need to add Stat on file here
	os.Mkdir(fmt.Sprintf("%s%s%s", UploadedFilePath, ImageDirectory, UID), 0766)
	ID := generateID(UID, name)
	err = ioutil.WriteFile(fmt.Sprintf("%s%s%s/%s", UploadedFilePath, ImageDirectory, UID, ID), data, 0666)

	if err != nil {
		return jsonResponseErr(err, 500)
	}

	img := asset.NewImage(UID, ID)
	_, err = img.Store(mongo.Database(asset.DatabaseName))

	if err != nil {
		return jsonResponseErr(err, 500)
	}

	return jsonResponseOk(dataResponse{"id": img.ID, "url": asset.GenerateImageURL(asset.TmpImageViewingPath, UID, ID)})
}

func deleteImage(id, uid string) ([]byte, int, error) {
	err := os.Remove(fmt.Sprintf("%s%s%s/%s", UploadedFilePath, ImageDirectory, uid, id))
	if err != nil {
		return jsonResponseErr(err, 500)
	}
	return jsonResponseOk(nil)
}
