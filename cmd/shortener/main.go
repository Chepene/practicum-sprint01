package main

import (
	"net/http"

	"github.com/Chepene/practicum-sprint01/internal/handler"
)

func main() {

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	mux := http.NewServeMux()
	mux.HandleFunc(`POST /`, handler.CreateHandler)
	mux.HandleFunc(`GET /{id}`, handler.GetByIDHandler)

	return http.ListenAndServe(`:8080`, mux)
}
