package service

import "github.com/Chepene/practicum-sprint01/internal/repository"

var baseUrl = "http://localhost:8080/"

func CreateShortLink(originalLink string) (string, error) {
	for {
		id := RandomString(8)
		targetLink := baseUrl + id
		if err := repository.SaveShortLink(originalLink, targetLink); err != nil {
			if err == repository.AlreadyExistsError {
				continue
			} else {
				return "", err
			}
		} else {
			return targetLink, nil
		}
	}
}
