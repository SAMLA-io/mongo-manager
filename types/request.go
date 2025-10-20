package types

import "go.mongodb.org/mongo-driver/v2/bson"

type Request struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
	Filter     bson.D `json:"filter,omitempty"`
}

type InsertOneRequest struct {
	Database   string                 `json:"database"`
	Collection string                 `json:"collection"`
	Data       map[string]interface{} `json:"data"`
}

type InsertManyRequest struct {
	Database   string                   `json:"database"`
	Collection string                   `json:"collection"`
	Data       []map[string]interface{} `json:"data"`
}
