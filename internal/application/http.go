package application

import (
	"github.com/edwsel/golang-common/pkg/runner/node"
	"github.com/labstack/echo/v4"
)

func (a *App) bootstrapHttp() error {
	http := node.NewHttpEcho(node.HttpEchoConfig{
		Host: a.config.Http.Host,
		Port: a.config.Http.Port,
	})

	a.runner.Add(http)

	http.GET("/time", func(c echo.Context) error {
		return nil
	})

	return nil
}
