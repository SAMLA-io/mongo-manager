package v1

import (
	"encoding/json"
	"mongo-manager/mongo"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, []string{"POST"}) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	request := GetRequest(r)

	docs, err := mongo.GetAll(request)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(docs)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, []string{"POST"}) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	request := GetRequest(r)

	doc, err := mongo.GetOne(request)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(doc)
}

func InsertOne(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, []string{"POST"}) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	request := GetInsertOneRequest(r)

	result, err := mongo.InsertOne(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func InsertMany(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, []string{"POST"}) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	request := GetInsertManyRequest(r)

	result, err := mongo.InsertMany(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func UpdateOne(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, []string{"PUT"}) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	request := GetUpdateOneRequest(r)

	result, err := mongo.UpdateOne(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func UpdateMany(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, []string{"PUT"}) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	request := GetUpdateManyRequest(r)

	result, err := mongo.UpdateMany(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, []string{"DELETE"}) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	request := GetDeleteOneRequest(r)

	result, err := mongo.DeleteOne(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func DeleteMany(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, []string{"DELETE"}) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	request := GetDeleteManyRequest(r)

	result, err := mongo.DeleteMany(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
