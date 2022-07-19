package main

import (
	"carsharing/internal/Router"
	"carsharing/internal/config"
	"carsharing/internal/entity/car"
	"carsharing/internal/entity/user"
	"carsharing/pkg/logging"
	"carsharing/pkg/postgres"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(config.ProvideConfig, logging.ProvideLogger, Router.NewRouter),
		fx.Invoke(user.NewHandlerUser, car.NewHandlerCar, postgres.NewClient),
	).Run()
}
