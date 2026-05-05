package handler

import (
	"github.com/HosseinForouzan/workout-tracker.git/auth/authservice"
	"github.com/HosseinForouzan/workout-tracker.git/config"
	"github.com/HosseinForouzan/workout-tracker.git/user/service"
	"github.com/HosseinForouzan/workout-tracker.git/user/validator"
)

type Handler struct {
	config config.Config
	userSvc service.Service
	userValidator validator.Validator
	authSvc authservice.Service
}

func New(cfg config.Config ,userSvc service.Service, validator validator.Validator, authSvc authservice.Service) Handler {
	return Handler{
		config: cfg,
		userSvc: userSvc,
		userValidator: validator,
		authSvc: authSvc,
	}
}