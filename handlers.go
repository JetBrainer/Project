package main

import (
	"encoding/json"
	"github.com/GophersLang/GoRest/models"
	uuid "github.com/satori/go.uuid"
	"io"
	"net/http"
)
func handleGetTokens(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var account models.Account
		var tokens models.JWTPackage

		json.NewDecoder(r.Body).Decode(&account)
		tokens.Generate(account)
		respondWithJson(w,http.StatusOK,tokens)
	}
	c, err := r.Cookie("session")
	if err != nil{
		http.Redirect(w,r,"/", http.StatusSeeOther)
	}
	io.WriteString(w,c.String())
}

func handleUuid(w http.ResponseWriter, r *http.Request){
	c, err := r.Cookie("session")
	if err != nil{
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	io.WriteString(w,c.String())
}
