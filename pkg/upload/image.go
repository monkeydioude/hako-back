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
	ImageDirectory         = "img/"
	TmpUserId              = "0"
	TmpImageViewingBaseUrl = "http://localhost:8880"
)

func saveImage(file multipart.File, name string) ([]byte, int, error) {
	data, err := ioutil.ReadAll(file)

	if err != nil {
		return jsonResponseErr(err.Error(), 500)
	}

	// Might need to add Stat on file here
	os.Mkdir(fmt.Sprintf("%s%s%s", UploadedFilePath, ImageDirectory, TmpUserId), 0766)
	fileName, fileExtendedPath, fileURL := generateFileInfo(name)
	err = ioutil.WriteFile(fmt.Sprintf("%s%s", UploadedFilePath, fileExtendedPath), data, 0666)

	if err != nil {
		return jsonResponseErr(err.Error(), 500)
	}

	img := asset.NewImage(fileName, TmpUserId, fileURL)
	_, err = img.Store(mongo.Database(asset.DatabaseName))

	if err != nil {
		return jsonResponseErr(err.Error(), 500)
	}

	return jsonResponseOk(dataResponse{"id": img.ID, "url": fileURL})
}
