package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// MongoClient is a struct for Mongo client
type MongoClient struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// NewClient creates a new instance of MongoClient with default settings.
func NewClient(uri string, dbName string) (*MongoClient, error) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db := client.Database(dbName)
	return &MongoClient{Client: client, DB: db}, nil
}

// InsertOne inserts a single document into a collection
func (m *MongoClient) InsertOne(collection string, document interface{}) (*mongo.InsertOneResult, error) {
	return m.DB.Collection(collection).InsertOne(context.TODO(), document)
}

// FindOne finds a single document in a collection
func (m *MongoClient) FindOne(collection string, filter bson.M) *mongo.SingleResult {
	return m.DB.Collection(collection).FindOne(context.TODO(), filter)
}

// FindAll finds all documents in a collection
func (m *MongoClient) FindAll(collection string, filter bson.M) (*mongo.Cursor, error) {
	return m.DB.Collection(collection).Find(context.TODO(), filter)
}

// UpdateOne updates a single document in a collection
func (m *MongoClient) UpdateOne(collection string, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	return m.DB.Collection(collection).UpdateOne(context.TODO(), filter, update)
}

// DeleteOne deletes a single document in a collection
func (m *MongoClient) DeleteOne(collection string, filter bson.M) (*mongo.DeleteResult, error) {
	return m.DB.Collection(collection).DeleteOne(context.TODO(), filter)
}

// UpsertOne updates a single document in a collection
func (m *MongoClient) UpsertOne(collection string, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	opts := options.Update().SetUpsert(true)
	return m.DB.Collection(collection).UpdateOne(context.TODO(), filter, update, opts)
}

// Close closes the Mongo connection
func (m *MongoClient) Close() error {
	return m.Client.Disconnect(context.TODO())
}
