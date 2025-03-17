package main

import (
	"BankRaya/config"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main(){
	config.ConnectDB()

	port := "8080"

	router := httprouter.New()

	log.Printf("Server running on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))	
}