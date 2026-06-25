package handler

import (
	"github.com/HosseinForouzan/workout-tracker.git/config"
	"github.com/HosseinForouzan/workout-tracker.git/exercise/service"
)

type Handler struct {
	config config.Config
	exerciseSvc service.Service
}

func New(cfg config.Config, exerciseSvc service.Service) Handler {
	return Handler{
		config: cfg,
		exerciseSvc: exerciseSvc,
	}
}