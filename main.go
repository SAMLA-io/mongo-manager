package main

import (
	"log"
	v1 "mongo-manager/api/v1"
	"mongo-manager/auth"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/health", healthCheck)

	// V1 API

	http.Handle("/v1/get-all", auth.VerifyingMiddleware(http.HandlerFunc(v1.GetAll)))
	http.Handle("/v1/get-one", auth.VerifyingMiddleware(http.HandlerFunc(v1.GetOne)))
	http.Handle("/v1/insert-one", auth.VerifyingMiddleware(http.HandlerFunc(v1.InsertOne)))
	http.Handle("/v1/insert-many", auth.VerifyingMiddleware(http.HandlerFunc(v1.InsertMany)))
	http.Handle("/v1/update-one", auth.VerifyingMiddleware(http.HandlerFunc(v1.UpdateOne)))
	http.Handle("/v1/update-many", auth.VerifyingMiddleware(http.HandlerFunc(v1.UpdateMany)))
	http.Handle("/v1/delete-one", auth.VerifyingMiddleware(http.HandlerFunc(v1.DeleteOne)))
	http.Handle("/v1/delete-many", auth.VerifyingMiddleware(http.HandlerFunc(v1.DeleteMany)))

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      http.DefaultServeMux,
	}

	log.Fatal(server.ListenAndServe())
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
