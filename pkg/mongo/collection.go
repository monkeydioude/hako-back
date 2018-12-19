package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type collection struct {
	Name string
	coll *mongo.Collection
}

func (c *collection) InsertOne(data interface{}) (interface{}, error) {
	ctx, cf := context.WithTimeout(context.Background(), QueryTimeout)
	res, err := c.coll.InsertOne(ctx, data)

	cf()
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}
