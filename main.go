package main

import (
	"fmt"
	"log"
	"mongo-manager/mongo"
	"mongo-manager/types"
)

func main() {
	docs, err := mongo.GetAll(types.Request{
		Database:   "general",
		Collection: "nucleus",
		Data:       nil,
	})

	if err != nil {
		log.Fatalf("Error getting documents: %v", err)
	}
	fmt.Println(docs)
}
