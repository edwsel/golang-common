package action

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/a-system/golang-common/internal/provider"
	"time"
)

type TimeActon struct {
	provider *provider.Provider
}

func NewTimeActon(provider *provider.Provider) *TimeActon {
	return &TimeActon{provider: provider}
}

func (a *TimeActon) Handler(ctx *fiber.Ctx) error {
	a.provider.Log.Info("use TimeAction")

	return ctx.JSON(fiber.Map{
		"now": time.Now().Format("2006-01-02 15:04:05.000000"),
	})
}
