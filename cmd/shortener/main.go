package main

import (
	"net/http"

	"github.com/Chepene/practicum-sprint01/internal/app"
)

func main() {

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	app, err := app.NewApp()

	if err != nil {
		return err
	}

	return http.ListenAndServe(`:8080`, app.Mux)
}
