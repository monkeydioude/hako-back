package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/monkeydioude/hako-back/pkg/asset"
	"github.com/monkeydioude/hako-back/pkg/mongo"
	"github.com/monkeydioude/moon"
)

const (
	connectTimeout = 10 * time.Second
)

var serverPort string
var mongoDBAddr string

func init() {
	serverPort = os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("[ERR ] valid SERVER_PORT env var must be given")
	}

	mongoDBAddr = os.Getenv("MONGODB_ADDR")
	if mongoDBAddr == "" {
		log.Fatal("[ERR ] valid MONGODB_ADDR env var must be given")
	}
}

func main() {
	mongo.Connect(mongoDBAddr, connectTimeout)
	server := moon.Moon()
	server.AddHeader("Access-Control-Allow-Origin", "*")

	server.MakeRouter(
		moon.Get("/image/all", asset.GetAllImage),
		moon.Delete("/image/{id}", asset.DeleteImage),
	)

	moon.ServerRun(fmt.Sprintf(":%s", serverPort), server)
}
