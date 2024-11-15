package web

import (
	"github.com/labstack/echo/v4"
)

type ServerConfig struct {
}

func SetupServer() *echo.Echo {
	e := echo.New()

	return e
}
