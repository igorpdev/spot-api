package domain

func CalculateScore(place *Place, hasLocation bool, lat, lng float64, requestedProfiles []string) float64 {
	score := 0.0

	if hasLocation {
		dist := Distance(lat, lng, place.Lat, place.Lng)
		proximityScore := 1.0 / (1.0 + dist/10.0)
		score += proximityScore * 0.5
	}

	if len(requestedProfiles) > 0 {
		matchedProfiles := 0
		for _, reqProfile := range requestedProfiles {
			for _, placeProfile := range place.Profiles {
				if string(placeProfile) == reqProfile {
					matchedProfiles++
					break
				}
			}
		}
		if matchedProfiles > 0 {
			profileScore := float64(matchedProfiles) / float64(len(requestedProfiles))
			score += profileScore * 0.5
		}
	}

	return score
}
