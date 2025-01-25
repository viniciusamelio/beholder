package main

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/services"
	"beholder-api/internal/utils"
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	services.InitSomClient()

	ds := services.NewSomDatasource()
	repository := repositories.NewEnvironmentRepository(*ds)
	sessionRepo := repositories.NewSessionRepository(*ds)
	callRepo := repositories.NewCallRepository(*ds)

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
					path := "/user"
					queryParams := "org_id=ashduashdu&something=127129"
					body := "{'name' : 'Lorem ipsum', 'email' : 'lorem@ipsum.com'}"
					callRepo.Create(models.Call{
						Name:        "Create User",
						Method:      "POST",
						QueryParams: &queryParams,
						Body:        &body,
						CalledAt:    time.Now(),
						SessionUID:  &e.UID,
						Session:     e,
						Path:        &path,
					}).Fold(
						func(f utils.Failure) {
							fmt.Printf("Failure: %s\n", f.Message())
						},
						func(e *models.Call) {
							fmt.Printf("Call created successfully: %d \n", e.UID)

						})
				},
			)
		},
	)

}
