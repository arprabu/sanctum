package repository

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var Items []Item

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)

	id, err := strconv.Atoi(parameter["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, item := range Items {
		if item.Id == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.NotFound(w, r)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item.Id = len(Items) + 1
	Items = append(Items, item)

	json.NewEncoder(w).Encode(item)
}
