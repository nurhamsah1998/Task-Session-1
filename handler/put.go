package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-session-1/models"
	"task-session-1/storage"
)

func UpdateCategories(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var payload models.Category
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	for i, category := range storage.Categories {
		if category.ID == id {
			storage.Categories[i].Name = payload.Name
			storage.Categories[i].Description = payload.Description

			json.NewEncoder(w).Encode(storage.Categories[i])
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]any{
		"success": false,
		"message": "Category not found",
	})
}
