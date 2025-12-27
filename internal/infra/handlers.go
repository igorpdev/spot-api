package infra

import (
	"encoding/json"
	"net/http"
	"spot-api/internal/domain"
	"strconv"
)

type Handlers struct {
	placeRepo domain.PlaceRepository
	cityRepo  domain.CityRepository
}

func NewHandlers(placeRepo domain.PlaceRepository, cityRepo domain.CityRepository) *Handlers {
	return &Handlers{placeRepo: placeRepo, cityRepo: cityRepo}
}

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func (h *Handlers) GetPlaces(w http.ResponseWriter, r *http.Request) {
	citySlug := r.URL.Query().Get("city")
	if citySlug == "" {
		http.Error(w, "city é obrigatório", http.StatusBadRequest)
		return
	}

	city, err := h.cityRepo.FindBySlug(citySlug)
	if err != nil {
		http.Error(w, "cidade não encontrada", http.StatusNotFound)
		return
	}

	places, err := h.placeRepo.FindByCity(city.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	latStr := r.URL.Query().Get("lat")
	lngStr := r.URL.Query().Get("lng")
	radiusStr := r.URL.Query().Get("radius")
	_ = r.URL.Query().Get("profiles")

	if latStr != "" {
		_, err := strconv.ParseFloat(latStr, 64)
		if err != nil {
			http.Error(w, "lat inválido", http.StatusBadRequest)
			return
		}
	}

	if lngStr != "" {
		_, err := strconv.ParseFloat(lngStr, 64)
		if err != nil {
			http.Error(w, "lng inválido", http.StatusBadRequest)
			return
		}
	}

	if radiusStr != "" {
		_, err := strconv.ParseFloat(radiusStr, 64)
		if err != nil {
			http.Error(w, "radius inválido", http.StatusBadRequest)
			return
		}
	}

	responses := make([]PlaceResponse, len(places))
	for i, place := range places {
		responses[i] = ToPlaceResponse(place, 0)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

func (h *Handlers) GetPlaceBySlug(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		http.Error(w, "slug é obrigatório", http.StatusBadRequest)
		return
	}

	place, err := h.placeRepo.FindBySlug(slug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := ToPlaceResponse(place, 0)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
