package mongo

import (
	"context"
	"fmt"

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

func (c *Collection) Find(filter interface{}) (*Cursor, error) {
	ctx, cf := context.WithTimeout(context.Background(), QueryTimeout)
	cur, err := c.Collection.Find(ctx, filter)
	defer cf()

	if err != nil {
		return nil, err
	}

	return &Cursor{
		Cursor: cur,
		Ctx:    context.Background(),
	}, nil
}

func (c *Collection) FindOne(filter interface{}) (*Element, error) {
	ctx, cf := context.WithTimeout(context.Background(), QueryTimeout)
	sr := c.Collection.FindOne(ctx, filter)
	defer cf()

	return &Element{
		Result: sr,
		Ctx:    context.Background(),
	}, nil
}

func (c *Collection) DeleteOne(filter interface{}) error {
	ctx, cf := context.WithTimeout(context.Background(), QueryTimeout)
	defer cf()
	dr, err := c.Collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if dr.DeletedCount == 0 {
		return fmt.Errorf("Could not delete resource from database")
	}

	return nil
}
