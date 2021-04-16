package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("api/users", GetUsers).Methods("GET")
	r.HandleFunc("api/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("api/users", CreateUser).Methods("POST")
	r.HandleFunc("api/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("api/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}
