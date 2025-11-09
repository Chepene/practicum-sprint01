package repository

import "sync"

type InMemoryLinkRepository struct {
	mu    sync.RWMutex
	links map[string]string
}

func NewInMemoryLinkRepository() *InMemoryLinkRepository {
	return &InMemoryLinkRepository{
		links: make(map[string]string),
	}
}

func (r *InMemoryLinkRepository) Save(originalLink string, targetLink string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.links[targetLink]; ok {
		return ErrAlreadyExists
	}

	r.links[targetLink] = originalLink

	return nil
}

func (r *InMemoryLinkRepository) GetOriginalLink(shortLink string) (string, error) {
	r.mu.Lock()
	r.mu.Unlock()

	originalLink, ok := r.links[shortLink]
	if !ok {
		return "", ErrNotFound
	}
	return originalLink, nil
}
