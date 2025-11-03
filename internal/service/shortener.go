package service

import "github.com/Chepene/practicum-sprint01/internal/repository"

type ShortenerService interface {
	CreateShortLink(originalLink string) (string, error)
	GetOriginalLink(shortLink string) (string, error)
}

type ShortenerServiceImpl struct {
	repo    repository.LinkRepository
	g       StringGenerator
	baseURL string
}

func NewShortenerService(repo repository.LinkRepository, g StringGenerator, baseURL string) *ShortenerServiceImpl {
	return &ShortenerServiceImpl{
		repo:    repo,
		g:       g,
		baseURL: baseURL,
	}
}

func (s *ShortenerServiceImpl) CreateShortLink(originalLink string) (string, error) {
	for {
		len := 8
		id := s.g.Generate(len)
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
