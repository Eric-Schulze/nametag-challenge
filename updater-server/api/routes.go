package api

import (
	"github.com/go-chi/chi/v5"
)

func (handler *UpdaterHandler) RegisterRoutes(router chi.Router) {
	router.Route("/updater", func(router chi.Router) {
		router.Get("/latest", handler.HandleLatestVersionNumber())
		router.Get("/checksum", handler.HandleGetCurrentChecksum())
		router.Post("/upload", handler.HandleUpload())
	})
}