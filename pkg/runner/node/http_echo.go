package node

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
)

type HttpEchoConfig struct {
	Host string
	Port uint
}

type HttpEcho struct {
	*echo.Echo
	Config HttpEchoConfig
}

func NewHttpEcho(config HttpEchoConfig) *HttpEcho {
	return &HttpEcho{
		Echo:   echo.New(),
		Config: config,
	}
}

func (h *HttpEcho) Name() string {
	return "http.echo"
}

func (h *HttpEcho) Run() error {
	return h.Echo.Start(fmt.Sprintf("%s:%d", h.Config.Host, h.Config.Port))
}

func (h *HttpEcho) Close() error {
	return h.Shutdown(context.Background())
}
