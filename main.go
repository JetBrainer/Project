package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handleUuid).Methods("POST")
	r.HandleFunc("/", handleUuid).Methods("GET")
	r.HandleFunc("/getTokens",handleGetTokens).Methods("POST")
	r.HandleFunc("/getTokens",handleGetTokens).Methods("GET")
	fmt.Println("Listening on :8080 port")

	r.HandleFunc("/us", InitCrud().ListUsers).Methods("GET")
	r.HandleFunc("/users", InitCrud().ListUsers).Methods("GET")
	r.HandleFunc("/users/new", InitCrud().AddForm).Methods("GET")
	r.HandleFunc("/users/new", InitCrud().AddUsers).Methods("POST")
	r.HandleFunc("/users/{id}", InitCrud().EditUsers).Methods("GET")
	r.HandleFunc("/users/{id}", InitCrud().UpdateUsers).Methods("POST")
	r.HandleFunc("/users/{id}", InitCrud().DeleteUsers).Methods("DELETE")

	r.HandleFunc("/tokens", InitCrud().ListSessions).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
