package monitorHandlers

import (
	"github.com/MarkTBSS/go-logger/config"
	"github.com/MarkTBSS/go-logger/modules/entities"
	"github.com/MarkTBSS/go-logger/modules/monitor"
	"github.com/gofiber/fiber/v2"
)

func MonitorHandler(cfg config.IConfig) IMontitorHandler {
	return &monitorHandler{
		cfg: cfg,
	}
}

type monitorHandler struct {
	cfg config.IConfig
}

type IMontitorHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

func (h *monitorHandler) HealthCheck(c *fiber.Ctx) error {
	res := &monitor.Monitor{
		Url:     h.cfg.App().Url(),
		Name:    h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}
	return entities.NewResponse(c).Success(fiber.StatusOK, res).Res()
}
