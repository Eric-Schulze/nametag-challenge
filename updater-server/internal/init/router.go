package init

import (
	"net/http"
	"updater-server/api"
	"updater-server/internal/common/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func initRouter(logger logger.Logger, updaterHandler api.UpdaterHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Heartbeat("/ping"))
	
	router.Handle("/data/*", http.StripPrefix("/data/world-pop/", http.FileServer(http.Dir("./data"))))


	if updaterHandler == (api.UpdaterHandler{}) {
		logger.Error("updaterHandler does not implement UpdaterHandler interface")
	} else {
		updaterHandler.RegisterRoutes(router)
	}

	return router
}
