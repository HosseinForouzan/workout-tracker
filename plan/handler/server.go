package handler

import (
	"github.com/HosseinForouzan/workout-tracker.git/config"
	"github.com/HosseinForouzan/workout-tracker.git/plan/service"
)

type Handler struct {
	config config.Config
	planSvc service.Service
}

func New(cfg config.Config, planSvc service.Service) Handler {
	return Handler{
		config: cfg,
		planSvc: planSvc,
	}
}