package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task-session-1/models"
	"task-session-1/services"
)

type CategoryHandler struct {
	service services.CategoryService
}

func NerCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	category, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "success",
		"data":    category,
	})
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	category, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "success",
		"data":    category,
	})
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	if err := h.service.Create(&category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "successfully create category",
	})
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	category.ID = id
	if err := h.service.Update(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "successfully update category",
	})
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "successfully delete category",
	})
}
