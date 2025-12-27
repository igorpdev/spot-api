package infra

import (
	"errors"
	"spot-api/internal/domain"
	"sync"
)

type MemoryRepository struct {
	places map[string]*domain.Place
	mu     sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		places: make(map[string]*domain.Place),
	}
}

func (r *MemoryRepository) FindAll() ([]*domain.Place, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	places := make([]*domain.Place, 0, len(r.places))
	for _, place := range r.places {
		places = append(places, place)
	}
	return places, nil
}

func (r *MemoryRepository) FindBySlug(slug string) (*domain.Place, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	place, exists := r.places[slug]
	if !exists {
		return nil, errors.New("lugar não encontrado")
	}
	return place, nil
}

func (r *MemoryRepository) Save(place *domain.Place) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := place.Validate(); err != nil {
		return err
	}

	r.places[place.Slug] = place
	return nil
}
