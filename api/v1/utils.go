package v1

import (
	"encoding/json"
	"log"
	"mongo-manager/types"
	"net/http"
	"strings"
)

// VerifyMethod checks if the HTTP request method is in the list of allowed methods.
// This function is used to ensure endpoints only accept the correct HTTP methods.
//
// Parameters:
//   - r: HTTP request to verify
//   - allowedMethods: Array of allowed HTTP methods (e.g., ["GET", "POST"])
//
// Returns:
//   - bool: True if the request method is allowed, false otherwise
//
// Example:
//
//	VerifyMethod(r, []string{"POST"}) // Only allows POST
//	VerifyMethod(r, []string{"GET", "POST"}) // Allows both GET and POST
func VerifyMethod(r *http.Request, allowedMethods []string) bool {
	for _, method := range allowedMethods {
		if r.Method == strings.ToUpper(method) {
			return true
		}
	}
	return false
}

func GetRequest(r *http.Request) types.Request {

	database := r.URL.Query().Get("database")
	collection := r.URL.Query().Get("collection")

	var requestBody types.Request
	json.NewDecoder(r.Body).Decode(&requestBody)

	return types.Request{
		Database:   database,
		Collection: collection,
		Filter:     requestBody.Filter,
	}
}

func GetOneRequest(r *http.Request) types.Request {

	database := r.URL.Query().Get("database")
	collection := r.URL.Query().Get("collection")

	var requestBody types.Request
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return types.Request{}
	}

	if requestBody.Filter == nil {
		return types.Request{}
	}

	return types.Request{
		Database:   database,
		Collection: collection,
		Filter:     requestBody.Filter,
	}
}

func GetInsertOneRequest(r *http.Request) types.InsertOneRequest {

	database := r.URL.Query().Get("database")
	collection := r.URL.Query().Get("collection")

	var requestBody types.InsertOneRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return types.InsertOneRequest{}
	}

	return types.InsertOneRequest{
		Database:   database,
		Collection: collection,
		Data:       requestBody.Data,
	}
}

func GetInsertManyRequest(r *http.Request) types.InsertManyRequest {

	database := r.URL.Query().Get("database")
	collection := r.URL.Query().Get("collection")

	var requestBody types.InsertManyRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return types.InsertManyRequest{}
	}

	return types.InsertManyRequest{
		Database:   database,
		Collection: collection,
		Data:       requestBody.Data,
	}
}

func GetUpdateOneRequest(r *http.Request) types.UpdateOneRequest {

	database := r.URL.Query().Get("database")
	collection := r.URL.Query().Get("collection")
	objectId := r.URL.Query().Get("objectId")

	var requestBody types.UpdateOneRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return types.UpdateOneRequest{}
	}

	request := types.UpdateOneRequest{
		Database:   database,
		Collection: collection,
		ObjectId:   objectId,
		Data:       requestBody.Data,
	}

	log.Printf("Request: %+v", request)

	return request

}

func GetUpdateManyRequest(r *http.Request) types.UpdateManyRequest {

	database := r.URL.Query().Get("database")
	collection := r.URL.Query().Get("collection")

	var requestBody types.UpdateManyRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return types.UpdateManyRequest{}
	}

	return types.UpdateManyRequest{
		Database:   database,
		Collection: collection,
		Filter:     requestBody.Filter,
		Data:       requestBody.Data,
	}
}

func GetDeleteOneRequest(r *http.Request) types.DeleteOneRequest {

	database := r.URL.Query().Get("database")
	collection := r.URL.Query().Get("collection")
	objectId := r.URL.Query().Get("objectId")

	return types.DeleteOneRequest{
		Database:   database,
		Collection: collection,
		ObjectId:   objectId,
	}
}

func GetDeleteManyRequest(r *http.Request) types.DeleteManyRequest {

	database := r.URL.Query().Get("database")
	collection := r.URL.Query().Get("collection")

	var requestBody types.DeleteManyRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return types.DeleteManyRequest{}
	}

	return types.DeleteManyRequest{
		Database:   database,
		Collection: collection,
		Filter:     requestBody.Filter,
	}
}
