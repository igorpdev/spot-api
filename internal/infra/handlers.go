package infra

import (
	"encoding/json"
	"net/http"
	"sort"
	"spot-api/internal/domain"
	"strconv"
	"strings"
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
	profilesStr := r.URL.Query().Get("profiles")

	var lat, lng, radius float64
	var hasLocation bool

	if latStr != "" && lngStr != "" {
		var err error
		lat, err = strconv.ParseFloat(latStr, 64)
		if err != nil {
			http.Error(w, "lat inválido", http.StatusBadRequest)
			return
		}

		lng, err = strconv.ParseFloat(lngStr, 64)
		if err != nil {
			http.Error(w, "lng inválido", http.StatusBadRequest)
			return
		}

		hasLocation = true
		radius = 50.0
		if radiusStr != "" {
			radius, err = strconv.ParseFloat(radiusStr, 64)
			if err != nil {
				http.Error(w, "radius inválido", http.StatusBadRequest)
				return
			}
		}
	} else if radiusStr != "" {
		http.Error(w, "radius requer lat e lng", http.StatusBadRequest)
		return
	}

	var requestedProfiles []string
	if profilesStr != "" {
		requestedProfiles = strings.Split(profilesStr, ",")
		for i := range requestedProfiles {
			requestedProfiles[i] = strings.TrimSpace(requestedProfiles[i])
		}
	}

	var filteredPlaces []PlaceWithDistance

	for _, place := range places {
		dist, include := applyFilters(place, hasLocation, lat, lng, radius, requestedProfiles)
		if include {
			filteredPlaces = append(filteredPlaces, PlaceWithDistance{
				Place:    place,
				Distance: dist,
			})
		}
	}

	if hasLocation {
		sort.Slice(filteredPlaces, func(i, j int) bool {
			return filteredPlaces[i].Distance < filteredPlaces[j].Distance
		})
	}

	responses := make([]PlaceResponse, len(filteredPlaces))
	for i, pw := range filteredPlaces {
		responses[i] = ToPlaceResponse(pw.Place, pw.Distance)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

type PlaceWithDistance struct {
	Place    *domain.Place
	Distance float64
}

func applyFilters(place *domain.Place, hasLocation bool, lat, lng, radius float64, requestedProfiles []string) (float64, bool) {
	dist := 0.0

	if hasLocation {
		dist = domain.Distance(lat, lng, place.Lat, place.Lng)
		if dist > radius {
			return 0, false
		}
	}

	if len(requestedProfiles) > 0 {
		if !hasAllProfiles(place, requestedProfiles) {
			return 0, false
		}
	}

	return dist, true
}

func hasAllProfiles(place *domain.Place, requestedProfiles []string) bool {
	for _, reqProfile := range requestedProfiles {
		found := false
		for _, placeProfile := range place.Profiles {
			if string(placeProfile) == reqProfile {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
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
