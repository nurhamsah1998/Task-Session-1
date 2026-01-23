package handler

import (
	"encoding/json"
	"net/http"
	"task-session-1/models"
	"task-session-1/storage"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if len(storage.Categories) >= 5 {
		http.Error(w, "Limit has reached!!, you cannot create again", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	storage.LastID++
	category.ID = storage.LastID
	storage.Categories = append(storage.Categories, category)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
