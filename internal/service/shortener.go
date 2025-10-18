package service

import "github.com/Chepene/practicum-sprint01/internal/repository"

var baseUrl = "http://localhost:8080/"

func CreateShortLink(originalLink string) (string, error) {
	for {
		id := RandomString(8)
		if err := repository.SaveShortLink(originalLink, id); err != nil {
			if err == repository.AlreadyExistsError {
				continue
			} else {
				return "", err
			}
		} else {
			targetLink := baseUrl + id
			return targetLink, nil
		}
	}
}

func GetOriginalLink(shortLink string) (string, error) {
	return repository.GetOriginalLink(shortLink)
}
