package asset

import (
	"fmt"
	"net/http"

	"github.com/monkeydioude/hako-back/pkg/json"
	"github.com/monkeydioude/hako-back/pkg/mongo"
	"github.com/monkeydioude/moon"
)

const (
	DatabaseName = "asset"
)

func GetAllImage(r *moon.Request) ([]byte, int, error) {
	cur, err := mongo.Database(DatabaseName).Collection("asset").Find(
		mongo.FiltersFromURLValues(r.QueryString),
	)
	if err != nil {
		return []byte(`{
			"status": "could not find files for specific user",
			"code": 500
		}
		`), 500, nil
	}

	res := json.Array()

	err = cur.ForEach(func(c *mongo.Cursor) error {
		item := &Image{}
		c.Decode(item)
		item.URL = item.GenerateUrl(TmpImageViewingPath)
		res.Add(json.Marshal(item))
		return nil
	})

	if err != nil {
		return []byte(`{
			"status": "could not produce response",
			"code": 500,
		}
		`), 500, err
	}

	return res.Bytes(), 200, nil
}

func DeleteImage(r *moon.Request) ([]byte, int, error) {
	if _, ok := r.Matches["id"]; !ok {
		return []byte(`{
			"status": "not found",
			"code": 404
		}`), 404, nil
	}

	mColl := mongo.Database(DatabaseName).Collection("asset")

	res, err := mColl.FindAndDecodeOne(&Image{
		ID: r.Matches["id"],
	})

	if err != nil {
		return []byte(`{
			"status": "Could not find resource to delete",
			"code": 500,
		}
		`), 500, err
	}

	filters := mongo.FiltersFromURLValues(r.QueryString)
	filters["id"] = r.Matches["id"]
	filters["user_id"] = res.(*Image).GetUserID()

	err = mColl.DeleteOne(filters)

	if err != nil {
		return []byte(`{
			"status": "Could not delete resource",
			"code": 500,
		}
		`), 500, err
	}

	err = sendDeleteRequest(fmt.Sprintf(
		"%s/user/%s/image/%s",
		TmpUploadURL,
		filters["user_id"],
		filters["id"],
	))

	if err != nil {
		return []byte(`{
			"status": "Could not send delete request",
			"code": 500,
		}
		`), 500, err
	}

	return []byte(`{
		"status": "ok",
		"code": 200
	}`), 200, nil
}

func sendDeleteRequest(url string) error {
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", url, nil)
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}
