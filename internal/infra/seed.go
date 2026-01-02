package infra

import (
	"spot-api/internal/domain"
)

func SeedPlaces(placeRepo *MemoryRepository, cityRepo *MemoryCityRepository) error {
	spCity := &domain.City{
		ID:   "1",
		Slug: "sp",
		Name: "São Paulo",
	}

	if err := cityRepo.Save(spCity); err != nil {
		return err
	}

	places := []*domain.Place{
		{
			ID:     "1",
			Name:   "Bar do Zé",
			Slug:   domain.GenerateSlug("Bar do Zé"),
			CityID: "1",
			Lat:    -23.5505,
			Lng:    -46.6333,
			Profiles: []domain.Profile{
				domain.ProfileBoemio,
				domain.ProfileNoturno,
			},
			Description:        "boteco de esquina com cerveja gelada e petiscos simples",
			MakesSenseFor:      "quem quer descomprimir sem frescura depois do trabalho",
			DoesNotMakeSenseIf: "você quer impressionar alguém ou precisa de silêncio",
			Tags:               []string{"cerveja", "petiscos", "informal"},
		},
		{
			ID:     "2",
			Name:   "Mercado Municipal",
			Slug:   domain.GenerateSlug("Mercado Municipal"),
			CityID: "1",
			Lat:    -23.5431,
			Lng:    -46.6290,
			Profiles: []domain.Profile{
				domain.ProfileCaoticoUrbano,
				domain.ProfileBaixaGastro,
			},
		},
		{
			ID:     "3",
			Name:   "Pinacoteca",
			Slug:   domain.GenerateSlug("Pinacoteca"),
			CityID: "1",
			Lat:    -23.5336,
			Lng:    -46.6331,
			Profiles: []domain.Profile{
				domain.ProfileArtes,
				domain.ProfileContemplativo,
			},
			Description:        "museu de arte com acervo brasileiro e exposições temporárias",
			MakesSenseFor:      "quem quer ver arte com calma e sem pressa",
			DoesNotMakeSenseIf: "você não gosta de museus ou está com pressa",
			Tags:               []string{"museu", "arte", "cultura"},
		},
		{
			ID:     "4",
			Name:   "Avenida Paulista",
			Slug:   domain.GenerateSlug("Avenida Paulista"),
			CityID: "1",
			Lat:    -23.5614,
			Lng:    -46.6560,
			Profiles: []domain.Profile{
				domain.ProfileCaoticoUrbano,
				domain.ProfileContemplativo,
			},
		},
		{
			ID:     "5",
			Name:   "Parque Ibirapuera",
			Slug:   domain.GenerateSlug("Parque Ibirapuera"),
			CityID: "1",
			Lat:    -23.5874,
			Lng:    -46.6576,
			Profiles: []domain.Profile{
				domain.ProfileContemplativo,
			},
		},
		{
			ID:     "6",
			Name:   "Beco do Batman",
			Slug:   domain.GenerateSlug("Beco do Batman"),
			CityID: "1",
			Lat:    -23.5505,
			Lng:    -46.6912,
			Profiles: []domain.Profile{
				domain.ProfileArtes,
				domain.ProfileBoemio,
			},
		},
		{
			ID:     "7",
			Name:   "Liberdade",
			Slug:   domain.GenerateSlug("Liberdade"),
			CityID: "1",
			Lat:    -23.5596,
			Lng:    -46.6333,
			Profiles: []domain.Profile{
				domain.ProfileBaixaGastro,
				domain.ProfileCaoticoUrbano,
			},
		},
		{
			ID:     "8",
			Name:   "Vila Madalena",
			Slug:   domain.GenerateSlug("Vila Madalena"),
			CityID: "1",
			Lat:    -23.5489,
			Lng:    -46.6938,
			Profiles: []domain.Profile{
				domain.ProfileBoemio,
				domain.ProfileNoturno,
				domain.ProfileArtes,
			},
			Description:        "bairro boêmio com bares, galerias e grafites nas ruas",
			MakesSenseFor:      "quem quer sair à noite, ver arte urbana e beber com amigos",
			DoesNotMakeSenseIf: "você quer sossego ou não gosta de agito noturno",
			Tags:               []string{"bairro", "noite", "arte urbana", "bares"},
		},
		{
			ID:     "9",
			Name:   "Museu de Arte de São Paulo",
			Slug:   domain.GenerateSlug("Museu de Arte de São Paulo"),
			CityID: "1",
			Lat:    -23.5614,
			Lng:    -46.6560,
			Profiles: []domain.Profile{
				domain.ProfileArtes,
				domain.ProfileContemplativo,
			},
		},
		{
			ID:     "10",
			Name:   "Praça do Por do Sol",
			Slug:   domain.GenerateSlug("Praça do Por do Sol"),
			CityID: "1",
			Lat:    -23.5505,
			Lng:    -46.6912,
			Profiles: []domain.Profile{
				domain.ProfileContemplativo,
			},
		},
	}

	for _, place := range places {
		if err := placeRepo.Save(place); err != nil {
			return err
		}
	}

	return nil
}
