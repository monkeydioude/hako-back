package asset

import (
	"github.com/monkeydioude/hako-back/pkg/mongo"
)

const (
	UploadedFilePath       = "/tmp/upload/"
	ImageDirectory         = "img/"
	TmpUserId              = "0"
	TmpImageViewingBaseUrl = "http://localhost:8880"
)

type Asset interface {
	GetType() string
	GetURL() string
	GetID() string
	GetUserID() string
	GetDateCreation() int64
	mongo.Storeable
}

type AssetsResponse struct {
	Assets []Asset `json:"assets"`
}

func (ar *AssetsResponse) PushAsset(a Asset) {
	ar.Assets = append(ar.Assets, a)
}

func NewAssetsResponse() *AssetsResponse {
	return &AssetsResponse{}
}
