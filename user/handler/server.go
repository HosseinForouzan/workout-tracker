package handler

import (
	"github.com/HosseinForouzan/workout-tracker.git/user/service"
	"github.com/HosseinForouzan/workout-tracker.git/user/validator"
)

type Handler struct {
	userSvc service.Service
	userValidator validator.Validator
}

func New(userSvc service.Service, validator validator.Validator) Handler {
	return Handler{
		userSvc: userSvc,
		userValidator: validator,
	}
}