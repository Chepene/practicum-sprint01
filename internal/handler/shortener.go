package handler

import (
	"io"
	"net/http"

	"github.com/Chepene/practicum-sprint01/internal/service"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	targetLink, err := service.CreateShortLink(string(body))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(targetLink)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetByIDHandler(w http.ResponseWriter, r *http.Request) {

	shortLink := r.PathValue(`id`)

	originalLink, err := service.GetOriginalLink(shortLink)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add(`Location`, originalLink)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
