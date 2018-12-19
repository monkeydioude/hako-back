package main

import (
	"time"

	"github.com/monkeydioude/hako-back/pkg/asset"
	"github.com/monkeydioude/moon"
)

const (
	connectTimeout = 10 * time.Second
)

func main() {
	// mw.Connect("mongodb://localhost:27017", connectTimeout)

	a := moon.Moon(nil)
	a.WithHeader("Access-Control-Allow-Origin", "*")
	a.Routes.AddGet("image/all", asset.GetAllImage)
	moon.ServerRun(":8882", a)
}
