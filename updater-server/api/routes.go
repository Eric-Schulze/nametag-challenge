package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (handler *UpdaterHandler) RegisterRoutes(router chi.Router) {
	router.Route("/updater", func(router chi.Router) {
		router.Handle("/data/*", http.StripPrefix("/data/", http.FileServer(http.Dir("./data"))))

		router.Get("/latest", handler.HandleLatestVersionNumber())
	})
}