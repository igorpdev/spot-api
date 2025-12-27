package domain

type PlaceRepository interface {
	FindAll() ([]*Place, error)
	FindBySlug(slug string) (*Place, error)
	FindByCity(cityID string) ([]*Place, error)
	Save(place *Place) error
}
