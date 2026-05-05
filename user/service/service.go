package service

import (
	"context"

	"github.com/HosseinForouzan/workout-tracker.git/auth/authservice"
	"github.com/HosseinForouzan/workout-tracker.git/user/entity"
)

type Repository interface {
	Register(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	GetUserByID(ctx context.Context, userID uint) (entity.User, error)
}

type Service struct {
	repo Repository
	auth authservice.Service
}

func New(repo Repository, auth authservice.Service) Service {
	return Service{repo: repo, auth: auth}
}
