package infra

import (
	"errors"
	"spot-api/internal/domain"
	"sync"
)

type MemoryCityRepository struct {
	cities map[string]*domain.City
	mu     sync.RWMutex
}

func NewMemoryCityRepository() *MemoryCityRepository {
	return &MemoryCityRepository{
		cities: make(map[string]*domain.City),
	}
}

func (r *MemoryCityRepository) FindBySlug(slug string) (*domain.City, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, city := range r.cities {
		if city.Slug == slug {
			return city, nil
		}
	}
	return nil, errors.New("cidade não encontrada")
}

func (r *MemoryCityRepository) Save(city *domain.City) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.cities[city.ID] = city
	return nil
}
