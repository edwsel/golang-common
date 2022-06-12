package application

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"gitlab.com/a-system/golang-common/internal/action"
	"gitlab.com/a-system/golang-common/pkg/runner/node"
)

func (a *App) bootstrapHttp() error {
	http := node.NewHttpFiber(node.HttpEchoConfig{
		FiberConfig: fiber.Config{
			DisableStartupMessage: true,
			JSONEncoder:           json.Marshal,
			JSONDecoder:           json.Unmarshal,
			AppName:               a.config.Http.Name,
			ProxyHeader:           fiber.HeaderXForwardedFor,
		},
		Host: a.config.Http.Host,
		Port: a.config.Http.Port,
	})

	a.runner.Add(http)

	http.Use(etag.New())
	http.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: false,
	}))

	http.Get("/time", action.NewTimeActon(a.provider).Handler)

	return nil
}
