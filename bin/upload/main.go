package main

import (
	"time"

	"github.com/monkeydioude/hako-back/pkg/mongo"
	"github.com/monkeydioude/hako-back/pkg/upload"
	"github.com/monkeydioude/moon"
)

const (
	connectTimeout = 10 * time.Second
)

func optionsPoke(r *moon.Request, c *moon.Configuration) ([]byte, int, error) {
	return []byte("OK"), 200, nil
}

func main() {
	mongo.Connect("mongodb://localhost:27017", connectTimeout)

	u := moon.Moon(nil)
	u.WithHeader("Access-Control-Allow-Origin", "*")
	u.Routes.AddPost("upload/image", upload.Image)
	u.Routes.Add(".+", "OPTIONS", optionsPoke)
	moon.ServerRun(":8881", u)
}
