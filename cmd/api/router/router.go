package router

import (
	"beholder-api/internal/data/repositories"

	"github.com/labstack/echo/v4"
)

func Router(envRepo *repositories.EnvironmentRepository) {
	r := echo.New()
	EnvironmentRoutes(r, *envRepo)
	r.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	r.Logger.Fatal(r.Start(":8080"))
}
