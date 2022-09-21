package node

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
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
	err := h.Echo.Start(fmt.Sprintf("%s:%d", h.Config.Host, h.Config.Port))

	if err != nil && errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (h *HttpEcho) Close() error {
	return h.Shutdown(context.Background())
}
