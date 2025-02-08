package main

import (
	migrations "beholder-api"
	"beholder-api/cmd/api/router"
	"beholder-api/cmd/crons"
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/services"
	"beholder-api/internal/services/bucket"
	"beholder-api/internal/services/providers"
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	services.InitSomClient()
	r2 := bucket.InitR2Bucket()
	if os.Getenv("ENV") != "local" {
		r2.DownloadFile("backup", "beholder.db")
	}

	fx.New(
		fx.Provide(
			services.NewSomDatasource,
			services.InitSqlite,
			repositories.NewEnvironmentRepository,
			repositories.NewSessionRepository,
			providers.NewCloudTasksProvider,
			func() bucket.Bucket {
				return r2
			},
		),
		fx.Invoke(
			func(db *sql.DB) {
				goose.SetBaseFS(migrations.Migrations)
				goose.Up(db, "migrations")

			},
			crons.Crons,
			router.Router,
		),
	).Run()

}
