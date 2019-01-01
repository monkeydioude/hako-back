package asset

import (
	"github.com/monkeydioude/hako-back/pkg/mongo"
)

const (
	UploadedFilePath       = "/tmp/upload/"
	ImageDirectory         = "img/"
	TmpUserId              = "0"
	TmpImageViewingBaseUrl = "http://localhost:8880"
	TmpImageViewingPath    = "http://localhost:8880/img/"
	TmpUploadURL           = "http://localhost:8881/upload"
)

type Asset interface {
	GetType() string
	GetURL() string
	GetID() string
	GetUserID() string
	GetDateCreation() int64
	mongo.Storeable
}
