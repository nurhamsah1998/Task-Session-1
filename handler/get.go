package handler

import (
	"encoding/json"
	"net/http"
	"task-session-1/storage"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.Categories)
}
