package service

import (
	"BankRaya/entity"
	"BankRaya/repository"
	"encoding/json"
	"net/http"
	"strings"

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

func (is *ItemService) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var Item entity.Item
	err := json.NewDecoder(r.Body).Decode(&Item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Bad Request", "detail": "Invalid request body"})
		return
	}
	
	Item.ItemCode = p.ByName("id")
	err = is.ItemRepository.Update(Item)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"message": "Not Found", "detail": err.Error()})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "Internal Server Error", "detail": err.Error()})
		}
		return
	}

	response := map[string]interface{}{
		"message": "Success update",
		"shipment": Item,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (is *ItemService) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	Item, err := is.ItemRepository.Delete(p.ByName("id"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"message": "Not Found", "detail": err.Error()})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "Internal Server Error", "detail": err.Error()})
		}
		return
	}

	response := map[string]interface{}{
		"message": "Success delete",
		"shipment": Item,
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}