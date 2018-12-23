package mongo

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var client *mongo.Client

const (
	// QueryTimeout defines the max amount of time a query should take.
	// It should be used in context of every query to mongoDB
	QueryTimeout = 1 * time.Second
)

// Storeable interface defines that a struct
// is ready to store itself in the database.
// It is this struct responsability to know
// the Collection data should be stored in
type Storeable interface {
	Store(*DB) (interface{}, error)
}

// Spawnable interface defines that a struct
// can be used as a formater for decoding
type Spawnable interface {
	Spawn() Spawnable
}

// Connect start dialing with mongoDB and conncetes to it.
// Will fail and returns an error if timeout is reached.
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
