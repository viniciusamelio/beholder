package main

import (
	"beholder-api/cmd/api/router"
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/services"
	"beholder-api/internal/services/providers"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	services.InitSomClient()

	fx.New(
		fx.Provide(
			services.NewSomDatasource,
			services.InitSqlite,
			repositories.NewCallRepository,
			repositories.NewEnvironmentRepository,
			repositories.NewSessionRepository,
			providers.NewCloudTasksProvider,
		),
		fx.Invoke(
			router.Router,
		),
	).Run()

}
