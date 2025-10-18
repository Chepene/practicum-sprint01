package repository

import "errors"

var links = make(map[string]string)

var AlreadyExistsError = errors.New("Такой элемент уже есть")

func SaveShortLink(originalLink string, targetLink string) error {

	if _, ok := links[targetLink]; ok {
		return AlreadyExistsError
	}

	links[targetLink] = originalLink

	return nil
}

func GetOriginalLink(shortLink string) (string, error) {
	originalLink, ok := links[shortLink]
	if !ok {
		return "", errors.New("Ссылка не найдена")
	}
	return originalLink, nil
}
