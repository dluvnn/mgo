package mgo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection ...
type Collection struct {
	name   string
	db     string
	client string
	coll   *mongo.Collection
	// mutex   sync.Mutex
	timeout time.Duration
}

// GetAll ...
func (c *Collection) GetAll(x interface{}) error {
	v, err := c.Find(bson.M{})
	if err != nil {
		return err
	}
	return v.All(x)
}

// InsertOne inserts a single document into the collection.
func (c *Collection) InsertOne(document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	//c.mutex.Lock()
	//defer c.mutex.Unlock()
	defer cancel()
	return c.coll.InsertOne(ctx, document, opts...)
}

// InsertMany inserts the provided documents.
func (c *Collection) InsertMany(documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	//c.mutex.Lock()
	//defer c.mutex.Unlock()
	defer cancel()
	return c.coll.InsertMany(ctx, documents, opts...)
}

// UpdateOne updates a single document in the collection.
func (c *Collection) UpdateOne(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	//c.mutex.Lock()
	//defer c.mutex.Unlock()
	defer cancel()
	return c.coll.UpdateOne(ctx, filter, update, opts...)
}

// UpdateMany updates multiple documents in the collection.
func (c *Collection) UpdateMany(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	//c.mutex.Lock()
	//defer c.mutex.Unlock()
	defer cancel()
	return c.coll.UpdateMany(ctx, filter, update, opts...)
}

// Duplicate ...
func (c *Collection) Duplicate() (*Client, *Collection, error) {
	n := new(Collection)
	n.client = c.client
	n.db = c.db
	n.name = c.name
	n.timeout = c.timeout
	client, err := New(c.client)
	if err != nil {
		return nil, nil, err
	}
	db := client.Database(c.db)
	return client, db.Collection(c.name), nil
}

// Find then finds the documents matching a model.
func (c *Collection) Find(filter interface{}, opts ...*options.FindOptions) (*Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	//c.mutex.Lock()
	//defer c.mutex.Unlock()
	defer cancel()
	cursor, err := c.coll.Find(ctx, filter, opts...)
	if cursor != nil {
		return &Cursor{Cursor: cursor}, err
	}
	return nil, err
}

// FindEx duplicates and then finds the documents matching a model.
func (c *Collection) FindEx(filter interface{}, opts ...*options.FindOptions) (*Cursor, error) {
	client, collection, err := c.Duplicate()
	if client == nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), collection.timeout)
	defer cancel()
	defer client.Disconnect()
	cursor, err := collection.coll.Find(ctx, filter, opts...)
	if cursor != nil {
		return &Cursor{Cursor: cursor}, err
	}
	return nil, err
}

// FindOne returns up to one document that matches the model.
func (c *Collection) FindOne(filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	//c.mutex.Lock()
	//defer c.mutex.Unlock()
	defer cancel()
	return c.coll.FindOne(ctx, filter, opts...)
}

// FindOneAndDelete find a single document and deletes it, returning the original in result.
func (c *Collection) FindOneAndDelete(filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	//c.mutex.Lock()
	//defer c.mutex.Unlock()
	defer cancel()
	return c.coll.FindOneAndDelete(ctx, filter, opts...)
}

// FindOneAndUpdate finds a single document and updates it, returning either the original or the updated.
func (c *Collection) FindOneAndUpdate(filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	//c.mutex.Lock()
	//defer c.mutex.Unlock()
	defer cancel()
	return c.coll.FindOneAndUpdate(ctx, filter, update, opts...)
}

// DeleteOne deletes a single document from the collection.
func (c *Collection) DeleteOne(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	//c.mutex.Lock()
	//defer c.mutex.Unlock()
	defer cancel()
	return c.coll.DeleteOne(ctx, filter, opts...)
}

// DeleteMany deletes multiple documents from the collection.
func (c *Collection) DeleteMany(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	//c.mutex.Lock()
	//defer c.mutex.Unlock()
	defer cancel()
	return c.coll.DeleteMany(ctx, filter, opts...)
}

// ReplaceOne replaces a single document in the collection.
func (c *Collection) ReplaceOne(filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	//c.mutex.Lock()
	//defer c.mutex.Unlock()
	defer cancel()
	return c.coll.ReplaceOne(ctx, filter, replacement, opts...)
}

// Indexes ..
func (c *Collection) Indexes(models []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()
	return c.coll.Indexes().CreateMany(ctx, models, opts...)
}

// CountDocuments ...
func (c *Collection) CountDocuments(filter interface{}, opts ...*options.CountOptions) (int64, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()
	return c.coll.CountDocuments(ctx, filter, opts...)
}

// EstimatedDocumentCount ...
func (c *Collection) EstimatedDocumentCount(opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()
	return c.coll.EstimatedDocumentCount(ctx, opts...)
}
