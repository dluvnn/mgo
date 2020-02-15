package mgo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client ...
type Client struct {
	*mongo.Client
	name    string
	timeout time.Duration
}

// Database ...
func (c *Client) Database(s string, opts ...*options.DatabaseOptions) *DB {
	db := new(DB)
	db.Database = c.Client.Database(s, opts...)
	db.timeout = c.timeout
	db.name = s
	db.client = c.name
	return db
}

// Ping ...
func (c *Client) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	return c.Client.Ping(ctx, readpref.Primary())
}

// Disconnect ...
func (c *Client) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	return c.Client.Disconnect(ctx)
}

// New opens new connection to mongo server
func New(uri string, timeout ...time.Duration) (*Client, error) {
	t := 30 * time.Second
	if len(timeout) == 1 {
		t = timeout[0]
	}
	client := new(Client)
	client.name = uri
	client.timeout = t
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	var err error
	client.Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	return client, client.Ping()
}
