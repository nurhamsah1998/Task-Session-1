package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"task-session-1/database"
	"task-session-1/handlers"
	"task-session-1/repositories"
	"task-session-1/services"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	DBConnection string `mapstructure:"DB_CONNECTION"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}
	config := Config{
		Port:         viper.GetString("PORT"),
		DBConnection: viper.GetString("DB_CONNECTION"),
	}

	// Setup database
	db, err := database.InitDB(config.DBConnection)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NerCategoryHandler(categoryService)

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			categoryHandler.GetAll(w, r)
		case http.MethodPost:
			categoryHandler.Create(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodPut:
			categoryHandler.Update(w, r)
		case http.MethodDelete:
			categoryHandler.Delete(w, r)
		case http.MethodGet:
			categoryHandler.GetByID(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal()
	} else {
		log.Println("🚀 Server running on http://localhost:8080")
	}

}
