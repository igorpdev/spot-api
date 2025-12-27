package infra

import (
	"encoding/json"
	"net/http"
	"spot-api/internal/domain"
)

type Handlers struct {
	repo domain.PlaceRepository
}

func NewHandlers(repo domain.PlaceRepository) *Handlers {
	return &Handlers{repo: repo}
}

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func (h *Handlers) GetPlaces(w http.ResponseWriter, r *http.Request) {
	places, err := h.repo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(places)
}

func (h *Handlers) GetPlaceBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		http.Error(w, "slug é obrigatório", http.StatusBadRequest)
		return
	}

	place, err := h.repo.FindBySlug(slug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(place)
}
