package router

import (
	"github.com/labstack/echo/v4"
)

func ErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, map[string]string{
		"error": message,
	})
}

func Response(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, data)
}
