package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type Collection struct {
	Name       string
	Collection *mongo.Collection
}

func (c *Collection) InsertOne(data interface{}) (interface{}, error) {
	ctx, cf := context.WithTimeout(context.Background(), QueryTimeout)
	res, err := c.Collection.InsertOne(ctx, data)
	defer cf()

	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

func (c *Collection) Find(data interface{}) (*Cursor, error) {
	ctx, cf := context.WithTimeout(context.Background(), QueryTimeout)
	cur, err := c.Collection.Find(ctx, data)
	defer cf()

	if err != nil {
		return nil, err
	}

	return &Cursor{
		Cursor: cur,
		Ctx:    context.Background(),
	}, nil
}

func (c *Collection) FindOne(data interface{}) (*Element, error) {
	ctx, cf := context.WithTimeout(context.Background(), QueryTimeout)
	sr := c.Collection.FindOne(ctx, data)
	defer cf()

	return &Element{
		Result: sr,
		Ctx:    context.Background(),
	}, nil
}
