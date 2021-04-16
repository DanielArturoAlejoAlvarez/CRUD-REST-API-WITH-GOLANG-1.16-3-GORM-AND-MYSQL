package main

import "github.com/gorilla/mux"

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("api/users", GetUsers).Methods("GET")
	r.HandleFunc("api/user/{id}", GetUser).Methods("GET")

}
