package asset

type Asset struct {
	Type string `json:"type"`
	URL  string `json:"url"`
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
