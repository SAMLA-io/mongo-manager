package api

import (
	"encoding/json"
	"mongo-manager/mongo"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	if !VerifyMethod(r, []string{"GET"}) {
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
	if !VerifyMethod(r, []string{"GET"}) {
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
