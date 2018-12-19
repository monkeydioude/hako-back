package asset

const (
	UploadedFilePath       = "/tmp/upload/"
	ImageDirectory         = "img/"
	TmpUserId              = "0"
	TmpImageViewingBaseUrl = "http://localhost:8880"
)

type Asset struct {
	Type string `bson:"type" json:"type"`
	URL  string `bson:"url" json:"url"`
}

type AssetsResponse struct {
	Assets []Asset `json:"assets"`
}

func (ar *AssetsResponse) PushAsset(t, url string) {
	ar.Assets = append(ar.Assets, Asset{
		Type: t,
		URL:  url,
	})
}

func NewAssetsResponse() *AssetsResponse {
	return &AssetsResponse{}
}
