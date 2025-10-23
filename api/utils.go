package api

import (
	"encoding/json"
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
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return types.Request{}
	}

	request := types.Request{
		Database:   database,
		Collection: collection,
		Filter:     requestBody.Filter,
	}

	return request
}
