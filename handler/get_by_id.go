package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task-session-1/storage"
)

func GetCategoriesById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// contoh path: /categories/1
	parts := strings.Split(r.URL.Path, "/")
	// ["", "categories", "1"]

	if len(parts) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Invalid URL",
		})
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil || id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	for i, category := range storage.Categories {
		if category.ID == id {

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
