package main

import (
	"net/http"
	"os"

	"github.com/monkeydioude/hako-back/pkg/upload"
	"github.com/monkeydioude/moon"
)

func optionsPoke(r *moon.Request, c *moon.Configuration) ([]byte, int, error) {
	return []byte("OK"), 200, nil
}

func main() {
	h := moon.Moon(nil)
	h.WithHeader("Access-Control-Allow-Origin", "*")

	os.Mkdir(upload.UploadedFilePath, 0766)
	os.Mkdir(upload.UploadedFilePath+upload.ImageDirectory, 0766)
	http.Handle("/", http.FileServer(http.Dir(upload.UploadedFilePath)))
	go http.ListenAndServe(":8880", nil)

	h.Routes.AddPost("upload/image", upload.Image)
	h.Routes.Add(".+", "OPTIONS", optionsPoke)
	moon.ServerRun(":8881", h)
}
