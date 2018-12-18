package main

import (
	"net/http"
	"os"

	"github.com/monkeydioude/hako-back/pkg/asset"
	"github.com/monkeydioude/hako-back/pkg/upload"
	"github.com/monkeydioude/moon"
)

func optionsPoke(r *moon.Request, c *moon.Configuration) ([]byte, int, error) {
	return []byte("OK"), 200, nil
}

func main() {
	// file service
	os.Mkdir(upload.UploadedFilePath, 0766)
	os.Mkdir(upload.UploadedFilePath+upload.ImageDirectory, 0766)
	http.Handle("/", http.FileServer(http.Dir(upload.UploadedFilePath)))
	go http.ListenAndServe(":8880", nil)

	// upload service
	u := moon.Moon(nil)
	u.WithHeader("Access-Control-Allow-Origin", "*")
	u.Routes.AddPost("upload/image", upload.Image)
	u.Routes.Add(".+", "OPTIONS", optionsPoke)
	go moon.ServerRun(":8881", u)

	// asset service
	a := moon.Moon(nil)
	a.WithHeader("Access-Control-Allow-Origin", "*")
	a.Routes.AddGet("image/all", asset.GetAllImage)
	moon.ServerRun(":8882", a)
}
