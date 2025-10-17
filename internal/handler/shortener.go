package handler

import (
	"io"
	"net/http"
	"strings"

	"github.com/Chepene/practicum-sprint01/internal/service"
)

var links = make(map[string]string)

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	targetLink, err := service.CreateShortLink(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(targetLink)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetByIDHandler(w http.ResponseWriter, r *http.Request) {

	var id = r.PathValue(`id`)

	var link, ok = links[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Добавляем протокол, если его нет
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = "http://" + link
	}

	w.Header().Add(`Location`, link)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
