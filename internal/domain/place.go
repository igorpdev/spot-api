package domain

import (
	"errors"
	"strings"
)

type Profile string

const (
	ProfileBoemio        Profile = "boemio"
	ProfileBaixaGastro   Profile = "baixa_gastro"
	ProfileAltaGastro    Profile = "alta_gastro"
	ProfileArtes         Profile = "artes"
	ProfileContemplativo Profile = "contemplativo"
	ProfileNoturno       Profile = "noturno"
	ProfileCaoticoUrbano Profile = "caotico_urbano"
)

type Place struct {
	ID                 string
	Name               string
	Slug               string
	CityID             string
	Lat                float64
	Lng                float64
	Profiles           []Profile
	Description        string
	MakesSenseFor      string
	DoesNotMakeSenseIf string
	Tags               []string
}

func (p *Place) Validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("nome não pode estar vazio")
	}

	if p.CityID == "" {
		return errors.New("cityID é obrigatório")
	}

	if len(p.Profiles) == 0 {
		return errors.New("pelo menos um perfil é obrigatório")
	}

	if p.Lat < -90 || p.Lat > 90 {
		return errors.New("latitude inválida")
	}

	if p.Lng < -180 || p.Lng > 180 {
		return errors.New("longitude inválida")
	}

	return nil
}

// GenerateSlug gera um slug simples a partir do nome
// Pode ser tosco, melhora depois
func GenerateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "ã", "a")
	slug = strings.ReplaceAll(slug, "á", "a")
	slug = strings.ReplaceAll(slug, "à", "a")
	slug = strings.ReplaceAll(slug, "â", "a")
	slug = strings.ReplaceAll(slug, "é", "e")
	slug = strings.ReplaceAll(slug, "ê", "e")
	slug = strings.ReplaceAll(slug, "í", "i")
	slug = strings.ReplaceAll(slug, "ó", "o")
	slug = strings.ReplaceAll(slug, "ô", "o")
	slug = strings.ReplaceAll(slug, "õ", "o")
	slug = strings.ReplaceAll(slug, "ú", "u")
	slug = strings.ReplaceAll(slug, "ç", "c")
	return slug
}
