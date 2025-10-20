package mongo

import (
	"context"
	"log"
	"mongo-manager/types"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var Client *mongo.Client

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)

	var err error
	Client, err = mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
}

func GetAll(request types.Request) ([]bson.M, error) {
	collection := Client.Database(request.Database).Collection(request.Collection)

	// Ensure we never pass a nil top-level filter; MongoDB requires a document, not null
	filter := request.Filter
	if filter == nil {
		filter = bson.D{}
	}
	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Printf("Error finding documents: %v", err)
		return nil, err
	}

	var docs []bson.M
	if err := cursor.All(context.TODO(), &docs); err != nil {
		log.Printf("Error decoding documents: %v", err)
		return nil, err
	}
	return docs, nil
}

func GetOne(request types.Request) (bson.M, error) {
	collection := Client.Database(request.Database).Collection(request.Collection)

	filter := request.Filter
	if filter == nil {
		filter = bson.D{}
	}

	doc := bson.M{}
	err := collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		log.Printf("Error finding document: %v", err)
		return nil, err
	}
	return doc, nil
}

func InsertOne(request types.InsertOneRequest) (*mongo.InsertOneResult, error) {
	collection := Client.Database(request.Database).Collection(request.Collection)

	result, err := collection.InsertOne(context.TODO(), request.Data)
	if err != nil {
		log.Printf("Error inserting document: %v", err)
		return nil, err
	}
	return result, nil
}

func InsertMany(request types.InsertManyRequest) (*mongo.InsertManyResult, error) {
	collection := Client.Database(request.Database).Collection(request.Collection)

	result, err := collection.InsertMany(context.TODO(), request.Data)
	if err != nil {
		log.Printf("Error inserting documents: %v", err)
		return nil, err
	}
	return result, nil
}
