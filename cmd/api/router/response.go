package router

import (
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
)

func ErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, map[string]string{
		"error": message,
	})
}

func Response(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, data)
}

func ProtobufResponse(c echo.Context, code int, data proto.Message) error {
	res, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	return c.Blob(code, "application/x-protobuf", res)
}
