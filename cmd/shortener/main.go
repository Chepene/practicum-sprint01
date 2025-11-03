package main

import (
	"net/http"

	"github.com/Chepene/practicum-sprint01/internal/handler"
	"github.com/Chepene/practicum-sprint01/internal/pkg"
	"github.com/Chepene/practicum-sprint01/internal/repository"
	"github.com/Chepene/practicum-sprint01/internal/service"
)

func main() {

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	repo := repository.NewInMemoryLinkRepository()

	g := pkg.NewRandomGenerator()

	baseUrl := "http://localhost:8082/"
	service := service.NewShortenerService(repo, g, baseUrl)

	handler := handler.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc(`POST /`, handler.CreateHandler)
	mux.HandleFunc(`GET /{id}`, handler.GetByIDHandler)

	return http.ListenAndServe(`:8082`, mux)
}
