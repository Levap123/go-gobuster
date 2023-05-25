package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func (h *Handlers) InitRoutes() http.Handler {
	chi := chi.NewMux()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Здесь можно указать список разрешенных источников
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Применяем middleware CORS ко всем маршрутам
	chi.Use(cors.Handler)
	chi.Get("/bust", h.handleBust)

	return chi
}
