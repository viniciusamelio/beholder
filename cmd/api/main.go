package main

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/services"
	"beholder-api/internal/utils"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	services.InitSomClient()

	ds := services.NewSomDatasource()
	repository := repositories.NewEnvironmentRepository(*ds)
	sessionRepo := repositories.NewSessionRepository(*ds)

	repository.Create(models.Environment{
		Name:    "local-env",
		BaseUrl: "https://local.env.com",
		Tags:    []string{"local", "dev", "coo-tag"},
	}).Fold(
		func(f utils.Failure) {
			fmt.Printf("Failure: %s\n", f.Message())
		},
		func(created *models.Environment) {
			fmt.Printf("Success: %d	\n", created.UID)
			sessionRepo.Create(&models.Session{
				UserID: "my-user-id",
				EnvUID: created.UID,
				Env:    *created,
			}).Fold(
				func(f utils.Failure) {
					fmt.Printf("Failure: %s\n", f.Message())
				},
				func(e *models.Session) {
					fmt.Printf("Success: %d	\n", e.UID)
					sessionRepo.Delete(e.UID).Fold(
						func(f utils.Failure) {
							fmt.Printf("Failure: %s\n", f.Message())
						},
						func(e bool) {
							fmt.Printf("Session deleted")
						},
					)
				},
			)
		},
	)

}
