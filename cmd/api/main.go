package main

import (
	"encoding/json"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "notification-service/docs"
	"notification-service/internal/config"
	"notification-service/internal/database"
	"notification-service/internal/notification"
)

type HealthResponse struct {
	Status string `json:"status"`
}

// @title Notification Service API
// @version 1.0
// @description API for the Notification Service
// @host localhost:8080
// @BasePath /
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgres(cfg.DB.DSN())

	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&notification.Notification{}); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	repository := notification.NewRepository(db)
	notificationHandler := notification.NewHandler(repository)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(HealthResponse{
			Status: "ok",
		})
	})

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	mux.HandleFunc("GET /notifications", notificationHandler.List)
	mux.HandleFunc("POST /notifications", notificationHandler.Create)

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: mux,
	}

	log.Println("Server started on port", cfg.AppPort)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
