package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Cursor ...
type Cursor struct {
	*mongo.Cursor
}

// Next gets the next result from this cursor.
// Returns true if there were no errors and the next result is available for decoding.
func (c *Cursor) Next() bool {
	return c.Cursor.Next(context.TODO())
}

func (c *Cursor) Close() error {
	return c.Cursor.Close(context.TODO())
}

func (c *Cursor) All(result interface{}) error {
	return c.Cursor.All(context.TODO(), result)
}
