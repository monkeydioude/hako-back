package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/monkeydioude/hako-back/pkg/asset"
	"github.com/monkeydioude/moon"
)

const (
	connectTimeout = 10 * time.Second
)

var serverPort string

func init() {
	serverPort = os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("[ERR ] valid SERVER_PORT env var must be given")
	}
}

func main() {
	a := moon.Moon(nil)
	a.WithHeader("Access-Control-Allow-Origin", "*")
	a.Routes.AddGet("image/all", asset.GetAllImage)
	moon.ServerRun(fmt.Sprintf(":%s", serverPort), a)
}
