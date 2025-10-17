package repository

import "errors"

var baseUrl = "http://localhost:8080"

var links = make(map[string]string)

var AlreadyExistsError = errors.New("Такой элемент уже есть")

func SaveShortLink(originalLink string, targetLink string) error {

	if _, ok := links[targetLink]; ok {
		return AlreadyExistsError
	}

	links[targetLink] = originalLink

	return nil
}
