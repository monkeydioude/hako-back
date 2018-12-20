package upload

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/monkeydioude/tools"
)

const (
	UploadedFilePath = "/tmp/upload/"
)

func getFileName(userID, name string) (string, string) {
	p := strings.Split(name, ".")
	m := tools.RandUnixNano(100)
	ext := ""
	fileName := tools.MD5(fmt.Sprintf("%s%s%d%d", TmpUserId, name, tools.RandUnixNano(m*100), time.Now().Unix())).String()
	if len(p) >= 2 {
		ext = fmt.Sprintf(".%s", p[1])
	}

	return fileName, ext
}

func generateFileInfo(name string) (string, string, string) {
	fileName, fileNameExt := getFileName(TmpUserId, name)
	fileExtendedPath := fmt.Sprintf("%s%s/%s%s", ImageDirectory, TmpUserId, fileName, fileNameExt)
	fileURL := fmt.Sprintf("%s/%s", TmpImageViewingBaseUrl, fileExtendedPath)

	return fileName, fileExtendedPath, fileURL
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

func jsonResponseOk(data dataResponse) ([]byte, int, error) {
	return jsonResponse(response{
		Status: "ok",
		Code:   200,
		Data:   data,
	}), 200, nil
}

func jsonResponseErr(status string, code int) ([]byte, int, error) {
	return jsonResponse(response{
		Status: status,
		Code:   int16(code),
	}), code, nil
}
