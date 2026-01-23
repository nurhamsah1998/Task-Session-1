package main

import (
	"log"
	"net/http"
	"task-session-1/handler"
)

func main() {
	http.HandleFunc("/category", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			handler.GetCategories(w, r)
		case http.MethodPost:
			handler.CreateCategory(w, r)
		case http.MethodPut:
			handler.UpdateCategories(w, r)
		case http.MethodDelete:
			handler.DeleteCategories(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
