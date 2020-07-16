package main

import (
	"fmt"
	"github.com/GophersLang/GoRest/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"html/template"
	"log"
	"net/http"
)

func InitCrud(){
	host := "mongodb://localhost:27017"
	dbname := "test"
	client, err := mgo.Dial(host)
	if err != nil{
		log.Fatal(err)
	}
	users := client.DB(dbname).C("user")
	sessions := client.DB(dbname).C("session")
	handlers := &models.Handler{
		Users: users,
		Sessions: sessions,
		Tmpl:  template.Must(template.ParseGlob("./template/*")),
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.ListUsers).Methods("GET")
	r.HandleFunc("/users", handlers.ListUsers).Methods("GET")
	r.HandleFunc("/users/new", handlers.AddForm).Methods("GET")
	r.HandleFunc("/users/new", handlers.AddUsers).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.EditUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUsers).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.DeleteUsers).Methods("DELETE")

	r.HandleFunc("/tokens", handlers.ListSessions).Methods("GET")



	fmt.Println("localhost:8181")
	http.ListenAndServe(":8181", r)
}
