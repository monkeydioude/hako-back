package asset

import (
	"github.com/monkeydioude/hako-back/pkg/mongo"
	"github.com/monkeydioude/moon"
)

const (
	DatabaseName = "asset"
)

func getByUserID(userID string) ([]byte, int, error) {
	cur, err := mongo.Database(DatabaseName).Collection("asset").Find(&Image{
		UserID: userID,
	})
	if err != nil {
		return []byte(`{
			"status": "could not find files for specific user",
			"code": 500
		}
		`), 500, nil
	}

	res, err := cur.JSONMarshal(&Image{})

	if err != nil {
		return []byte(`{
			"status": "could not marshal files",
			"code": 500,
		}
		`), 500, err
	}

	return res, 200, nil
}

func GetAllImage(r *moon.Request) ([]byte, int, error) {
	if _, ok := r.QueryString["user_id"]; ok {
		return getByUserID(r.QueryString["user_id"])
	}

	return []byte(`{
		"status": "not found",
		"code": 404
	}`), 404, nil
}

func DeleteImage(r *moon.Request) ([]byte, int, error) {
	if _, ok := r.Matches["id"]; !ok {
		return []byte(`{
			"status": "not found",
			"code": 404
		}`), 404, nil
	}

	return []byte(`{
		"status": "ok",
		"code": 200
	}`), 200, nil
}
