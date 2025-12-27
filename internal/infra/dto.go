package infra

import "spot-api/internal/domain"

type PlaceResponse struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Slug     string   `json:"slug"`
	Lat      float64  `json:"lat"`
	Lng      float64  `json:"lng"`
	Profiles []string `json:"profiles"`
	Distance float64  `json:"distance,omitempty"`
}

func ToPlaceResponse(place *domain.Place, distance float64) PlaceResponse {
	profiles := make([]string, len(place.Profiles))
	for i, p := range place.Profiles {
		profiles[i] = string(p)
	}

	return PlaceResponse{
		ID:       place.ID,
		Name:     place.Name,
		Slug:     place.Slug,
		Lat:      place.Lat,
		Lng:      place.Lng,
		Profiles: profiles,
		Distance: distance,
	}
}
