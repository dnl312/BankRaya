package service

import (
	"BankRaya/entity"
	"BankRaya/repository"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ItemService struct {
	ItemRepository *repository.ItemRepository
}

func NewItemService(pr *repository.ItemRepository) *ItemService {
	return &ItemService{ItemRepository: pr}
}

func (is *ItemService) FindAll(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	data, err := is.ItemRepository.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal Server Error", "detail": err.Error()})	
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (is *ItemService) Insert(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	w.Header().Set("Content-Type", "application/json")

	var Item entity.Item
	err := json.NewDecoder(r.Body).Decode(&Item)

	if Item.ItemName == "" || Item.ItemPrice == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Bad Request", "detail": "ItemCode,ItemName, and ItemPrice are required"})
		return
	}

	err = is.ItemRepository.Insert(Item)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal Server Error", "detail": err.Error()})
		return
	}
	
	response := map[string]interface{}{
		"message": "Success create",
		"shipment": Item,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}