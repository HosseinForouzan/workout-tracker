package service

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/workout-tracker.git/user/param"
)

func (s Service) Profile(ctx context.Context, req param.ProfileRequest) (param.ProfileResponse, error) {
	user, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return param.ProfileResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.ProfileResponse{
		Name: user.Name,
		Email: user.Email,
	}, nil
}