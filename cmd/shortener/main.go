package main

import (
	"net/http"

	"github.com/Chepene/practicum-sprint01/internal/handler"
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

	baseUrl := "http://localhost:8080/"
	service := service.NewShortenerService(repo, baseUrl)

	handler := handler.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc(`POST /`, handler.CreateHandler)
	mux.HandleFunc(`GET /{id}`, handler.GetByIDHandler)

	return http.ListenAndServe(`:8080`, mux)
}
