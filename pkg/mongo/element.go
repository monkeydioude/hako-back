package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type Element struct {
	Result *mongo.DocumentResult
	Ctx    context.Context
}

// Decode is a gateway to mongo.Decode
func (e *Element) Decode(target interface{}) error {
	return e.Result.Decode(target)
}
