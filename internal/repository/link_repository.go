package repository

type InMemoryLinkRepository struct {
	links map[string]string
}

func NewInMemoryLinkRepository() *InMemoryLinkRepository {
	return &InMemoryLinkRepository{
		links: make(map[string]string),
	}
}

func (r *InMemoryLinkRepository) Save(originalLink string, targetLink string) error {

	if _, ok := r.links[targetLink]; ok {
		return ErrAlreadyExists
	}

	r.links[targetLink] = originalLink

	return nil
}

func (r *InMemoryLinkRepository) GetOriginalLink(shortLink string) (string, error) {
	originalLink, ok := r.links[shortLink]
	if !ok {
		return "", ErrNotFound
	}
	return originalLink, nil
}
