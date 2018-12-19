package main

import (
	"github.com/monkeydioude/hako-back/pkg/asset"
	"github.com/monkeydioude/moon"
)

func main() {
	a := moon.Moon(nil)
	a.WithHeader("Access-Control-Allow-Origin", "*")
	a.Routes.AddGet("image/all", asset.GetAllImage)
	moon.ServerRun(":8882", a)
}
