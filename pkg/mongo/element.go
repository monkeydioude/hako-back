package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type Element struct {
	Result *mongo.SingleResult
	Ctx    context.Context
}

// Decode is a gateway to mongo.Decode
func (e *Element) Decode(target interface{}) error {
	return e.Result.Decode(target)
}

func (c *Collection) FindAndDecodeOne(entity Spawnable) (Spawnable, error) {
	ele, err := c.FindOne(entity)

	if err != nil {
		return nil, err
	}

	res := entity.Spawn()
	err = ele.Decode(res)

	return res, err
}
