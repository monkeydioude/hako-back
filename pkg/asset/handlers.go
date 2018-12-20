package asset

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/monkeydioude/moon"
)

func getByUserId(userID string) ([]byte, int, error) {
	userImgPath := fmt.Sprintf("%s%s%s", UploadedFilePath, ImageDirectory, userID)
	files, err := ioutil.ReadDir(userImgPath)
	if err != nil {
		return []byte(`{
			"status": "could not find files for specific user",
			"code": 500
		}
		`), 500, nil
	}

	ar := NewAssetsResponse()
	for _, f := range files {
		ar.PushAsset(NewImage(f.Name(), userID, fmt.Sprintf("%s/%s%s/%s", TmpImageViewingBaseUrl, ImageDirectory, userID, f.Name())))
	}

	res, err := json.Marshal(ar)
	if err != nil {
		return []byte(`{
			"status": "could not marshal files",
			"code": 500
		}
		`), 500, nil
	}

	return res, 200, nil
}

func GetAllImage(r *moon.Request, c *moon.Configuration) ([]byte, int, error) {
	if _, ok := r.QueryString["user_id"]; ok {
		return getByUserId(r.QueryString["user_id"])
	}

	return []byte(`{
		"status": "not found",
		"code": 404
	}
	`), 404, nil
}
