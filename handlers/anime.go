package handlers

import (
	"net/http"
	"strconv"

	"anime-server/providers"

	"github.com/go-chi/chi/v5"
)

type AnimeHandler struct {
	Provider providers.AnimeProvider
}

func NewAnimeHandler(p providers.AnimeProvider) *AnimeHandler {
	return &AnimeHandler{Provider: p}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (h *AnimeHandler) SearchAnime(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		// http.Error(w, "Missing query param ?q=", http.StatusBadRequest)
		WriteError(w, http.StatusBadRequest, "Missing query param ?q=")
		return
	}

	page := 1
	limit := 5

	if p := r.URL.Query().Get("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil && val > 0 {
			page = val
		}
	}

	if l := r.URL.Query().Get("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil && val > 0 {
			limit = val
		}
	}

	data, err := h.Provider.Search(r.Context(), query, page, limit)
	if err != nil {
		// http.Error(w, "Failed to fetch anime", 500)
		WriteError(w, http.StatusInternalServerError, "Failed to fetch anime")
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(data)
	WriteSuccess(w, http.StatusOK, data)
}

func (h *AnimeHandler) GetAnimeByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// http.Error(w, "Invalid ID", http.StatusBadRequest)
		WriteError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	animeDetail, err := h.Provider.GetByID(r.Context(), id)
	if err != nil {
		// http.Error(w, "Failed to fetch anime details", 500)
		WriteError(w, http.StatusInternalServerError, "Failed to fetch anime details")
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(animeDetail)
	WriteSuccess(w, http.StatusOK, animeDetail)
}
