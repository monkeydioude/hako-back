package mongo

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

type DB struct {
	Name string
	db   *mongo.Database
}

func Database(name string) *DB {
	return &DB{
		Name: name,
		db:   client.Database(name),
	}
}

func (d *DB) Collection(name string) *Collection {
	return &Collection{
		Name: name,
		coll: d.db.Collection(name),
	}
}
