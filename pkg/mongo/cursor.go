package mongo

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// Cursor is, atm, just a gateway to mongo.Cursor
// until I find a clever way to wrap cleanly
// the iterate and decode pattern from Mongo
type Cursor struct {
	Cursor mongo.Cursor
	Ctx    context.Context
}

// Next is a gateway to mongo.Next
func (c *Cursor) Next() bool {
	return c.Cursor.Next(c.Ctx)
}

// Decode is a gateway to mongo.Decode
func (c *Cursor) Decode(target interface{}) error {
	return c.Cursor.Decode(target)
}

// Close is a gateway to mongo.Decode
func (c *Cursor) Close() error {
	return c.Cursor.Close(c.Ctx)
}

// ForEach process the chain of Cursors and call a function for every iteration
func (c *Cursor) ForEach(handler func(*Cursor) error) error {
	defer c.Close()
	for c.Cursor.Next(c.Ctx) {
		err := handler(c)
		if err != nil {
			return err
		}
	}
	return nil
}

// JSONMarshal decodes completly a Cursor, most likely after a Find(),
// and JsonMarshal it into a slice of byte, ready to be used as a response.
// If there is m
func (c *Cursor) JSONMarshal(data Spawnable) ([]byte, error) {
	it := 0
	buffer := bytes.NewBuffer(nil)

	for c.Next() {
		if it > 0 {
			_, err := buffer.WriteString(",")
			if err != nil {
				return nil, err
			}
		}
		data = data.Spawn()
		err := c.Decode(data)
		if err != nil {
			return nil, err
		}

		res, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		_, err = buffer.Write(res)
		if err != nil {
			return nil, err
		}
		it++
	}

	r := buffer.Bytes()
	buffer = bytes.NewBuffer([]byte("["))
	buffer.Write(r)
	buffer.WriteString("]")

	return buffer.Bytes(), nil
}
