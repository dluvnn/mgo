package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

/**
example:
err := c.Transaction(func(coll *mongo.Collection, ctx mongo.SessionContext) error {
	_, err := coll.UpdateOne(ctx, bson.M{"_id": id1}, update_1)
	if err != nil{
		return err
	}

	_, err := coll.UpdateOne(ctx, bson.M{"_id": id2}, update_2)
	if err != nil{
		return err
	}

	return nil
})
*/

// TransactionCallback ...
type TransactionCallback func(*mongo.Collection, mongo.SessionContext) error

// Transaction ...
func (c *Collection) Transaction(cb TransactionCallback) error {
	client, collection, err := c.Duplicate()
	if client == nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), collection.timeout)
	defer cancel()
	defer client.Disconnect()

	session, err := client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	if err = session.StartTransaction(); err != nil {
		return err
	}

	return mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err := cb(collection.coll, sc)
		if err != nil {
			return err
		}
		return session.CommitTransaction(sc)
	})
}
