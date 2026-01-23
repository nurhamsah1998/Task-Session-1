package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-session-1/storage"
)

func DeleteCategories(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
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
			storage.Categories = append(storage.Categories[:i], storage.Categories[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{
				"success": true,
				"message": "Category deleted successfully",
			})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]any{
		"success": false,
		"message": "Category not found",
	})
}
