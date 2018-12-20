package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/monkeydioude/hako-back/pkg/mongo"
	"github.com/monkeydioude/hako-back/pkg/upload"
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

func optionsPoke(r *moon.Request, c *moon.Configuration) ([]byte, int, error) {
	return []byte("OK"), 200, nil
}

func main() {
	mongo.Connect(mongoDBAddr, connectTimeout)

	u := moon.Moon(nil)
	u.WithHeader("Access-Control-Allow-Origin", "*")
	u.Routes.AddPost("upload/image", upload.Image)
	u.Routes.Add(".+", "OPTIONS", optionsPoke)
	moon.ServerRun(fmt.Sprintf(":%s", serverPort), u)
}
