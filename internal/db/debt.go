package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert(collection string, data any) (primitive.ObjectID, error) {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbName).Collection(collection)

	resp, err := c.InsertOne(context.Background(), data)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return resp.InsertedID.(primitive.ObjectID), nil

}

func FindByUser(collection string, document any) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbName).Collection(collection)

	cursor, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	return cursor.All(context.Background(), document)
}

func FindById(collection, id string, result interface{}) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbName).Collection(collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	return c.FindOne(context.Background(), filter).Decode(result)
}

func Find(collection string, document any) (*mongo.Cursor, error) {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbName).Collection(collection)

	cursor, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	return cursor, nil
}
