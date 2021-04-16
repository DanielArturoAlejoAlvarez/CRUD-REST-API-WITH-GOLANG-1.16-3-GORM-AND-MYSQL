package main

import "github.com/gorilla/mux"

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("api/users", GetUsers).Methods("GET")
	r.HandleFunc("api/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("api/users", CreateUser).Methods("POST")
}
