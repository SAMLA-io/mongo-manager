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

	if len(docs) == 0 {
		return []bson.M{}, nil
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

	if err == mongo.ErrNoDocuments {
		return bson.M{}, nil
	}

	if err != nil {
		log.Printf("Error finding document: %v", err)
		return bson.M{}, err
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

func UpdateOne(request types.UpdateOneRequest) (*mongo.UpdateResult, error) {
	collection := Client.Database(request.Database).Collection(request.Collection)

	objId, err := bson.ObjectIDFromHex(request.ObjectId)
	if err != nil {
		log.Printf("Error converting object ID: %v", err)
		return nil, err
	}
	update := bson.D{{Key: "$set", Value: request.Data}}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objId}, update)
	if err != nil {
		log.Printf("Error updating document: %v", err)
		return nil, err
	}
	return result, nil
}

func UpdateMany(request types.UpdateManyRequest) (*mongo.UpdateResult, error) {
	collection := Client.Database(request.Database).Collection(request.Collection)

	filter := request.Filter
	if filter == nil {
		filter = bson.D{}
	}
	update := bson.D{{Key: "$set", Value: request.Data}}

	result, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Printf("Error updating documents: %v", err)
		return nil, err
	}
	return result, nil
}

func DeleteOne(request types.DeleteOneRequest) (*mongo.DeleteResult, error) {
	collection := Client.Database(request.Database).Collection(request.Collection)

	objId, err := bson.ObjectIDFromHex(request.ObjectId)
	if err != nil {
		log.Printf("Error converting object ID: %v", err)
		return nil, err
	}

	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objId})
	if err != nil {
		log.Printf("Error deleting document: %v", err)
		return nil, err
	}
	return result, nil
}

func DeleteMany(request types.DeleteManyRequest) (*mongo.DeleteResult, error) {
	collection := Client.Database(request.Database).Collection(request.Collection)

	filter := request.Filter
	if filter == nil {
		filter = bson.D{}
	}

	// Convert string _id to ObjectID if present
	// Convert bson.D to bson.M for easier manipulation
	filterMap := bson.M{}
	for _, elem := range filter {
		filterMap[elem.Key] = elem.Value
	}

	if idValue, exists := filterMap["_id"]; exists {
		if idStr, ok := idValue.(string); ok {
			objId, err := bson.ObjectIDFromHex(idStr)
			if err != nil {
				log.Printf("Error converting object ID: %v", err)
				return nil, err
			}
			filterMap["_id"] = objId
		}
	}

	// Convert back to bson.D
	filter = bson.D{}
	for key, value := range filterMap {
		filter = append(filter, bson.E{Key: key, Value: value})
	}

	result, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Printf("Error deleting documents: %v", err)
		return nil, err
	}
	return result, nil
}
