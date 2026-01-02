package domain

import "testing"

func TestCalculateScore_WithLocationAndProfiles(t *testing.T) {
	place := &Place{
		ID:       "1",
		Name:     "Bar Teste",
		Lat:      -23.5505,
		Lng:      -46.6333,
		Profiles: []Profile{ProfileBoemio, ProfileNoturno},
	}

	hasLocation := true
	lat := -23.5505
	lng := -46.6333
	requestedProfiles := []string{"boemio"}

	score := CalculateScore(place, hasLocation, lat, lng, requestedProfiles)

	if score <= 0 {
		t.Errorf("esperado score > 0, obteve %f", score)
	}

	if score > 1.0 {
		t.Errorf("esperado score <= 1.0, obteve %f", score)
	}
}

func TestCalculateScore_OnlyProfiles(t *testing.T) {
	place := &Place{
		ID:       "1",
		Name:     "Bar Teste",
		Lat:      -23.5505,
		Lng:      -46.6333,
		Profiles: []Profile{ProfileBoemio, ProfileNoturno},
	}

	hasLocation := false
	lat := 0.0
	lng := 0.0
	requestedProfiles := []string{"boemio", "noturno"}

	score := CalculateScore(place, hasLocation, lat, lng, requestedProfiles)

	if score <= 0 {
		t.Errorf("esperado score > 0, obteve %f", score)
	}

	if score != 0.5 {
		t.Errorf("esperado score 0.5 (match completo de perfis), obteve %f", score)
	}
}

func TestCalculateReasons_WithLocationAndProfiles(t *testing.T) {
	place := &Place{
		ID:       "1",
		Name:     "Bar Teste",
		Lat:      -23.5505,
		Lng:      -46.6333,
		Profiles: []Profile{ProfileBoemio, ProfileNoturno},
	}

	hasLocation := true
	lat := -23.5505
	lng := -46.6333
	requestedProfiles := []string{"boemio"}

	reasons := CalculateReasons(place, hasLocation, lat, lng, requestedProfiles)

	if len(reasons) == 0 {
		t.Error("esperado pelo menos uma reason")
	}

	foundProximity := false
	foundProfile := false

	for _, reason := range reasons {
		if len(reason) > 6 && reason[:6] == "perto:" {
			foundProximity = true
		}
		if len(reason) > 7 && reason[:7] == "perfil:" {
			foundProfile = true
		}
	}

	if !foundProximity {
		t.Error("esperado reason de proximidade")
	}

	if !foundProfile {
		t.Error("esperado reason de perfil")
	}
}
