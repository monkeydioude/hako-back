package mongo

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

type database struct {
	Name string
	db   *mongo.Database
}

func Database(name string) *database {
	return &database{
		Name: name,
		db:   client.Database(name),
	}
}

func (d *database) Collection(name string) *collection {
	return &collection{
		Name: name,
		coll: d.db.Collection(name),
	}
}
