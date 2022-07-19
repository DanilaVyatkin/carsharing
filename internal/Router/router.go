package Router

import (
	"carsharing/internal/config"
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

func NewRouter(lc fx.Lifecycle, logger *zap.SugaredLogger, cfg *config.Config) *mux.Router {
	logger.Info("Executing NewRouter.")
	router := mux.NewRouter()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting HTTP server.")
			go http.ListenAndServe(cfg.ApplicationConfig.Address, router)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping HTTP server.")
			return logger.Sync()
		},
	})

	return router
}
