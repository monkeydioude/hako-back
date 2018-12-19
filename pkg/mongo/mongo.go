package mongo

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var client *mongo.Client

const (
	QueryTimeout = 5 * time.Second
)

func Connect(uri string, timeout time.Duration) error {
	var err error
	ctx, cf := context.WithTimeout(context.Background(), timeout)
	client, err = mongo.Connect(ctx, uri)

	cf()
	if err != nil {
		return err
	}

	return nil
}
