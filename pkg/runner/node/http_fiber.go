package node

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type HttpFiberConfig struct {
	FiberConfig fiber.Config
	Host        string
	Port        uint
}

type HttpFiber struct {
	*fiber.App
	Config HttpFiberConfig
}

func NewHttpFiber(config HttpFiberConfig) *HttpFiber {
	return &HttpFiber{
		App:    fiber.New(config.FiberConfig),
		Config: config,
	}
}

func (h *HttpFiber) Name() string {
	return "http.fiber"
}

func (h *HttpFiber) Run() error {
	return h.App.Listen(fmt.Sprintf("%s:%d", h.Config.Host, h.Config.Port))
}

func (h *HttpFiber) Close() error {
	return h.Shutdown()
}
