package upload

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/monkeydioude/tools"
)

const (
	UploadedFilePath = "/tmp/upload/"
)

func generateID(userID, name string) string {
	m := tools.RandUnixNano(100)
	now := time.Now().Unix()
	ID := tools.MD5(
		fmt.Sprintf(
			"%s%s%d%d",
			userID,
			name,
			tools.RandUnixNano(m*100),
			now,
		),
	).String()

	return fmt.Sprintf("%s%s%d", userID, ID, now)
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

func jsonResponseErr(err error, code int) ([]byte, int, error) {
	return jsonResponse(response{
		Status: err.Error(),
		Code:   int16(code),
	}), code, err
}

func jsonResponse404() ([]byte, int, error) {
	return jsonResponse(response{
		Status: "not found",
		Code:   404,
	}), 404, nil
}
