package main

import (
	"github.com/GophersLang/GoRest/models"
	"gopkg.in/mgo.v2"
	"html/template"
	"log"
)

func InitCrud() *models.Handler{
	host := "mongodb+srv://user1:<password>@cluster0.mru6s.mongodb.net/test?retryWrites=true&w=majority"
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
	return handlers

}
