package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"anime-server/handlers"
	"anime-server/providers/anilist"
	"anime-server/providers/jikan"

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

	//anilist
	anilistProvider := anilist.New(c)
	anilistHandler := handlers.NewAnimeHandler(anilistProvider)

	//jikan
	jikanProvider := jikan.New(c)
	jikanHandler := handlers.NewAnimeHandler(jikanProvider)
	// r.Get("/health", handlers.HealthCheck)
	r.Get("/anilist/anime/search", anilistHandler.SearchAnime)
	r.Get("/anilist/anime/{id}", anilistHandler.GetAnimeByID)
	r.Get("/jikan/anime/search", jikanHandler.SearchAnime)
	r.Get("/jikan/anime/{id}", jikanHandler.GetAnimeByID)

	server := &http.Server{Addr: ":8080", Handler: r}

	go func() {
		log.Println("Starting server on http://localhost:8080")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited properly")
}
