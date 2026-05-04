package service

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/workout-tracker.git/user/entity"
	"github.com/HosseinForouzan/workout-tracker.git/user/param"
)

func (s Service) Register(ctx context.Context, req param.RegisterReqeust) (param.RegisterResponse, error) {
	user := entity.User{
		ID: 0,
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
	}

	createdUser, err := s.repo.Register(ctx, user)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.RegisterResponse{
		Name: createdUser.Name,
	}, nil
}