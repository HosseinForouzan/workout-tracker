package service

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/workout-tracker.git/user/param"
)

func (s Service) Login(ctx context.Context, req param.LoginRequest) (param.LoginResponse, error) {

	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	if req.Email != user.Email || req.Password != user.Password {
		return param.LoginResponse{}, fmt.Errorf("user credential is wrong")
	}

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	refreshToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return param.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.LoginResponse{
		Name: user.Name,
		Email: user.Email,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}