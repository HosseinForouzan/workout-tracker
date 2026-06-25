package service

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/workout-tracker.git/exercise/entity"
	"github.com/HosseinForouzan/workout-tracker.git/exercise/param"
)

type Repository interface{
	CreateExercise(ctx context.Context, e entity.Exercise) error
	ListExercises(ctx context.Context) ([]entity.Exercise, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) AddExercise(ctx context.Context, req param.ExerciseRequest) error {
	exercise := entity.Exercise{
		ID: 0,
		Name: req.Name,
		MuscleGroup: req.MuscleGroup,
	}

	err := s.repo.CreateExercise(ctx, exercise)
	if err != nil {
		fmt.Errorf("unexpected error: %w", err)
	}

	return nil
}

func (s Service) ListExercise(ctx context.Context) ([]entity.Exercise, error) {
	exercises, err := s.repo.ListExercises(ctx)
	if err != nil {
		return nil, fmt.Errorf("unexpected error: %w", err)
	}


	return exercises, nil
}

