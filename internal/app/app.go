package app

import (
	"net/http"

	"github.com/Chepene/practicum-sprint01/internal/handler"
	"github.com/Chepene/practicum-sprint01/internal/pkg"
	"github.com/Chepene/practicum-sprint01/internal/repository"
	"github.com/Chepene/practicum-sprint01/internal/service"
)

type App struct {
	Mux     *http.ServeMux
	Handler *handler.Handler
}

func NewApp() (*App, error) {

	repo := repository.NewInMemoryLinkRepository()

	g := pkg.NewRandomGenerator()

	baseUrl := "http://localhost:8080/"
	service := service.NewShortenerService(repo, g, baseUrl)

	handler := handler.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc(`POST /`, handler.CreateHandler)
	mux.HandleFunc(`GET /{id}`, handler.GetByIDHandler)

	return &App{
		Mux:     mux,
		Handler: handler,
	}, nil
}
