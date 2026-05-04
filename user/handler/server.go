package handler

import "github.com/HosseinForouzan/workout-tracker.git/user/service"

type Handler struct {
	userSvc service.Service
}

func New(userSvc service.Service) Handler {
	return Handler{
		userSvc: userSvc,
	}
}