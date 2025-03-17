package service

import (
	"BankRaya/repository"
	"encoding/json"
	"net/http"
)

type ItemService struct {
	ItemRepository *repository.ItemRepository
}

func NewItemService(pr *repository.ItemRepository) *ItemService {
	return &ItemService{ItemRepository: pr}
}

func (ps *ItemService) FindAll(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	data, err := ps.ItemRepository.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal Server Error", "detail": err.Error()})	
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
