package main

import (
	"net/http"
	"os"

	"github.com/monkeydioude/hako-back/pkg/upload"
)

func main() {
	os.Mkdir(upload.UploadedFilePath, 0766)
	os.Mkdir(upload.UploadedFilePath+upload.ImageDirectory, 0766)
	http.Handle("/", http.FileServer(http.Dir(upload.UploadedFilePath)))
	http.ListenAndServe(":8880", nil)
}
