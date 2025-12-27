package infra

import (
	"spot-api/internal/domain"
)

func SeedPlaces(repo *MemoryRepository) error {
	places := []*domain.Place{
		{
			ID:   "1",
			Name: "Bar do Zé",
			Slug: domain.GenerateSlug("Bar do Zé"),
			Lat:  -23.5505,
			Lng:  -46.6333,
			Profiles: []domain.Profile{
				domain.ProfileBoemio,
				domain.ProfileNoturno,
			},
		},
		{
			ID:   "2",
			Name: "Mercado Municipal",
			Slug: domain.GenerateSlug("Mercado Municipal"),
			Lat:  -23.5431,
			Lng:  -46.6290,
			Profiles: []domain.Profile{
				domain.ProfileCaoticoUrbano,
				domain.ProfileBaixaGastro,
			},
		},
		{
			ID:   "3",
			Name: "Pinacoteca",
			Slug: domain.GenerateSlug("Pinacoteca"),
			Lat:  -23.5336,
			Lng:  -46.6331,
			Profiles: []domain.Profile{
				domain.ProfileArtes,
				domain.ProfileContemplativo,
			},
		},
		{
			ID:   "4",
			Name: "Avenida Paulista",
			Slug: domain.GenerateSlug("Avenida Paulista"),
			Lat:  -23.5614,
			Lng:  -46.6560,
			Profiles: []domain.Profile{
				domain.ProfileCaoticoUrbano,
				domain.ProfileContemplativo,
			},
		},
		{
			ID:   "5",
			Name: "Parque Ibirapuera",
			Slug: domain.GenerateSlug("Parque Ibirapuera"),
			Lat:  -23.5874,
			Lng:  -46.6576,
			Profiles: []domain.Profile{
				domain.ProfileContemplativo,
			},
		},
		{
			ID:   "6",
			Name: "Beco do Batman",
			Slug: domain.GenerateSlug("Beco do Batman"),
			Lat:  -23.5505,
			Lng:  -46.6912,
			Profiles: []domain.Profile{
				domain.ProfileArtes,
				domain.ProfileBoemio,
			},
		},
		{
			ID:   "7",
			Name: "Liberdade",
			Slug: domain.GenerateSlug("Liberdade"),
			Lat:  -23.5596,
			Lng:  -46.6333,
			Profiles: []domain.Profile{
				domain.ProfileBaixaGastro,
				domain.ProfileCaoticoUrbano,
			},
		},
		{
			ID:   "8",
			Name: "Vila Madalena",
			Slug: domain.GenerateSlug("Vila Madalena"),
			Lat:  -23.5489,
			Lng:  -46.6938,
			Profiles: []domain.Profile{
				domain.ProfileBoemio,
				domain.ProfileNoturno,
				domain.ProfileArtes,
			},
		},
		{
			ID:   "9",
			Name: "Museu de Arte de São Paulo",
			Slug: domain.GenerateSlug("Museu de Arte de São Paulo"),
			Lat:  -23.5614,
			Lng:  -46.6560,
			Profiles: []domain.Profile{
				domain.ProfileArtes,
				domain.ProfileContemplativo,
			},
		},
		{
			ID:   "10",
			Name: "Praça do Por do Sol",
			Slug: domain.GenerateSlug("Praça do Por do Sol"),
			Lat:  -23.5505,
			Lng:  -46.6912,
			Profiles: []domain.Profile{
				domain.ProfileContemplativo,
			},
		},
	}

	for _, place := range places {
		if err := repo.Save(place); err != nil {
			return err
		}
	}

	return nil
}
