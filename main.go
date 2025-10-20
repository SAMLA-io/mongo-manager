package main

import (
	"fmt"
	"log"
	"mongo-manager/mongo"
	"mongo-manager/types"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func main() {
	docs, err := mongo.GetAll(types.Request{
		Database:   "general",
		Collection: "nucleus",
		Filter:     bson.D{},
	})

	if err != nil {
		log.Fatalf("Error getting documents: %v", err)
	}
	fmt.Println(docs)
}
