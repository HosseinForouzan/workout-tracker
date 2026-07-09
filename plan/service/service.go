package service

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/workout-tracker.git/plan/entity"
	"github.com/HosseinForouzan/workout-tracker.git/plan/param"
)

type PlanRepository interface {
	Add(ctx context.Context, p entity.WorkoutPlan) (entity.WorkoutPlan, error)
	GetAll(ctx context.Context) ([]entity.WorkoutPlan, error)
	GetPlansOfUser(ctx context.Context, userID uint) ([]entity.WorkoutPlan, error)
}

type Service struct {
	planRepo PlanRepository
}

func New(planRepo PlanRepository) Service {
	return Service{
		planRepo: planRepo,
	}
}

func (s Service) AddPlan(ctx context.Context, req param.PlanAddRequest) (param.PlanAddResponse, error) {
	plan := entity.WorkoutPlan{
		ID: 0,
		UserID: req.UserID,
		Name: req.Name,
	}

	workoutPlan, err := s.planRepo.Add(ctx, plan)
	if err != nil {
		return param.PlanAddResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.PlanAddResponse{
		ID: workoutPlan.ID,
		UserID: workoutPlan.UserID,
		Name: workoutPlan.Name,
	}, nil
}

func (s Service) GetAllPlans(ctx context.Context) (param.PlansGetResponse, error) {
	workoutPlans, err := s.planRepo.GetAll(ctx)
	if err != nil {
		return param.PlansGetResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.PlansGetResponse{WorkoutPlans: workoutPlans}, nil
}

func(s Service) GetPlansOfUser(ctx context.Context, userID uint) (param.PlansOfUserGetResponse, error) {
	workoutPlans, err := s.planRepo.GetPlansOfUser(ctx, userID)
	if err != nil {
		return param.PlansOfUserGetResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.PlansOfUserGetResponse{WorkoutPlans: workoutPlans}, nil
}