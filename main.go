package main

import (
	"log"
	"net/http"

	"anime-server/handlers"
	"anime-server/providers/anilist"

	"anime-server/internal/cache"
	"anime-server/internal/middleware"

	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RateLimiter())
	c := cache.New()
	provider := anilist.New(c)
	animeHandler := handlers.NewAnimeHandler(provider)
	// r.Get("/health", handlers.HealthCheck)
	r.Get("/anime/search", animeHandler.SearchAnime)
	r.Get("/anime/{id}", animeHandler.GetAnimeByID)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
