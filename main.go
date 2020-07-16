package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleUuid)
	http.HandleFunc("/getTokens",handleGetTokens)
	fmt.Println("Listening on :8080 port")
	go InitCrud()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
