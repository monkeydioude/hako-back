package upload

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/monkeydioude/tools"
)

type response struct {
	Status string `json:"status"`
	Name   string `json:"name"`
	Code   int16  `json:"code"`
	Url    string `json:"url"`
}

const (
	UploadedFilePath = "/tmp/upload/"
)

func getFileName(userID, name string) string {
	p := strings.Split(name, ".")
	fileName := tools.MD5(fmt.Sprintf("%s%d", TmpUserId, time.Now().Unix())).String()
	if len(p) >= 2 {
		fileName = fmt.Sprintf("%s.%s", fileName, p[1])
	}

	return fileName
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
