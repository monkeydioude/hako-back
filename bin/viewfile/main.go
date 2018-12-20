package main

import (
	"log"
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

var serverPort string

func init() {
	serverPort = os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("[ERR ] valid SERVER_PORT env var must be given")
	}
}
