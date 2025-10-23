package repository

import "errors"

type LinkRepository interface {
	Save(originalLink string, targetLink string) error
	GetOriginalLink(shortLink string) (string, error)
}

var (
	ErrAlreadyExists = errors.New("Сущность уже существует")
	ErrNotFound      = errors.New("Сущность не найдена")
)
