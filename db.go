package mgo

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

// DB ...
type DB struct {
	*mongo.Database
	name    string
	client  string
	timeout time.Duration
}

// Collection ...
func (db *DB) Collection(s string, options ...*options.CollectionOptions) *Collection {
	cn := new(Collection)
	cn.coll = db.Database.Collection(s, options...)
	cn.db = db.name
	cn.client = db.client
	cn.timeout = db.timeout
	return cn
}
