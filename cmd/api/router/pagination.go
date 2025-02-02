package router

import (
	"beholder-api/internal/dtos"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaginationContext struct {
	echo.Context
	Take int64
	Skip int64
}

func PaginationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		skip, err := strconv.Atoi(c.QueryParam("skip"))
		if err != nil {
			skip = int(dtos.DefaultPagination.Skip.Int64)
		}
		take, err := strconv.Atoi(c.QueryParam("take"))
		if err != nil {
			take = int(dtos.DefaultPagination.Take.Int64)
		}

		return next(&PaginationContext{
			c,
			int64(take),
			int64(skip),
		})
	}
}
