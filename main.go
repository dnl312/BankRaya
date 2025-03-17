package main

import (
	"BankRaya/config"
	"BankRaya/repository"
	"BankRaya/service"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main(){
	DB := config.ConnectDB()

	port := "8080"

	router := httprouter.New()

	itemRepository := repository.NewItemRepository(DB)
	ItemService := service.NewItemService(itemRepository)

	router.GET("/items", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		ItemService.FindAll(w, r)
	})

	router.POST("/items-create", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ItemService.Insert(w, r, p)
	})


	log.Printf("Server running on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))	
}