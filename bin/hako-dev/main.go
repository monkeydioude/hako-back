package main

import (
	"net/http"
	"os"

	"github.com/monkeydioude/hako-back/pkg/upload"
	"github.com/monkeydioude/moon"
)

func main() {
	h := moon.Moon(nil)
	h.WithHeader("Access-Control-Allow-Origin", "*")

	os.Mkdir(upload.UploadedFilePath, 0766)
	os.Mkdir(upload.UploadedFilePath+upload.ImageDirectory, 0766)
	http.Handle("/", http.FileServer(http.Dir(upload.UploadedFilePath)))
	go http.ListenAndServe(":8880", nil)

	h.Routes.AddPost("upload/image", upload.Image)
	moon.ServerRun(":8881", h)
}
