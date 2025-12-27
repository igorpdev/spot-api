package domain

type CityRepository interface {
	FindBySlug(slug string) (*City, error)
	Save(city *City) error
}
