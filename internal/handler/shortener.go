package handler

import (
	"io"
	"net/http"

	"github.com/Chepene/practicum-sprint01/internal/service"
)

type Handler struct {
	service service.ShortenerService
}

func NewHandler(svc service.ShortenerService) *Handler {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	targetLink, err := h.service.CreateShortLink(string(body))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(targetLink)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) GetByIDHandler(w http.ResponseWriter, r *http.Request) {

	shortLink := r.PathValue(`id`)

	originalLink, err := h.service.GetOriginalLink(shortLink)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add(`Location`, originalLink)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
