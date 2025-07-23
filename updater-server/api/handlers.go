package api

import (
	"fmt"
	"net/http"
	"updater-server/internal/common/logger"
)

type UpdaterHandler struct {
	service *UpdaterService
	logger  *logger.Logger
}

func NewUpdaterHandler(logger *logger.Logger, dataFilePath string, dataFileName string) *UpdaterHandler {
	service := NewUpdaterService(dataFilePath, dataFileName, logger)

	return &UpdaterHandler{
		service: service,
		logger:  logger,
	}
}

func (handler *UpdaterHandler) HandleLatestVersionNumber() http.HandlerFunc {
	// Use this closure to prepare any objects needed for the handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.logger.Info("Handling request for latest version number")

		serviceName := r.URL.Query().Get("service-name")
		handler.logger.Info("Received service name from request", "serviceName", serviceName)

		version, err := handler.service.GetLatestVersionNumber(serviceName)
		if err != nil {
			handler.logger.Error("Error getting lastet version number", "error", err)
			http.Error(w, "Error getting lastet version number", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, version)
	})
}
