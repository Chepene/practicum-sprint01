package service

import "github.com/Chepene/practicum-sprint01/internal/repository"

type ShortenerService interface {
	CreateShortLink(originalLink string) (string, error)
	GetOriginalLink(shortLink string) (string, error)
}

type ShortenerServiceImpl struct {
	repo    repository.LinkRepository
	baseURL string
}

func NewShortenerService(repo repository.LinkRepository, baseURL string) *ShortenerServiceImpl {
	return &ShortenerServiceImpl{
		repo:    repo,
		baseURL: baseURL,
	}
}

func (s *ShortenerServiceImpl) CreateShortLink(originalLink string) (string, error) {
	for {
		id := RandomString(8)
		if err := s.repo.Save(originalLink, id); err != nil {
			if err != repository.ErrAlreadyExists {
				return "", err
			}
		} else {
			targetLink := s.baseURL + id
			return targetLink, nil
		}
	}
}

func (s *ShortenerServiceImpl) GetOriginalLink(shortLink string) (string, error) {
	return s.repo.GetOriginalLink(shortLink)
}
