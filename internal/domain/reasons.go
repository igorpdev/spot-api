package domain

import "fmt"

func CalculateReasons(place *Place, hasLocation bool, lat, lng float64, requestedProfiles []string) []string {
	var reasons []string

	if hasLocation {
		dist := Distance(lat, lng, place.Lat, place.Lng)
		distKm := fmt.Sprintf("%.1fkm", dist)
		reasons = append(reasons, "perto: "+distKm)
	}

	if len(requestedProfiles) > 0 {
		for _, reqProfile := range requestedProfiles {
			for _, placeProfile := range place.Profiles {
				if string(placeProfile) == reqProfile {
					reasons = append(reasons, "perfil: "+reqProfile)
					break
				}
			}
		}
	}

	return reasons
}
