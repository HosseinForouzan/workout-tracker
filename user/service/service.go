package service

import (
	"context"

	"github.com/HosseinForouzan/workout-tracker.git/user/entity"
)

type Repository interface {
	Register(ctx context.Context, user entity.User) (entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
