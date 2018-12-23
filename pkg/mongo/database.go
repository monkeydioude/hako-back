package mongo

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

type DB struct {
	Name string
	DB   *mongo.Database
}

func Database(name string) *DB {
	return &DB{
		Name: name,
		DB:   client.Database(name),
	}
}

func (d *DB) Collection(name string) *Collection {
	return &Collection{
		Name:       name,
		Collection: d.DB.Collection(name),
	}
}
