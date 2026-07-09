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

type PlanExerciseRepository interface{
	AddExercisePlan(ctx context.Context, p entity.PlanExercise) (entity.PlanExercise, error)
}

type Service struct {
	planRepo PlanRepository
	planExerciseRepo PlanExerciseRepository
}

func New(planRepo PlanRepository, planExerciseRepo PlanExerciseRepository) Service {
	return Service{
		planRepo: planRepo,
		planExerciseRepo: planExerciseRepo,
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

func(s Service) AddPlanExercise(ctx context.Context, req param.PlanExerciseRequest) (param.PlanExerciseResponse, error) {
	p := entity.PlanExercise{
		ID: 0,
		PlanID: req.PlanID,
		ExerciseID: req.ExerciseID,
		TargetSets: req.TargetSets,
		TargetReps: req.TargetReps,
		TargetWeight: float64(req.TargetWeight),
		Order: req.SortOrder,
	}

	planExercise, err := s.planExerciseRepo.AddExercisePlan(ctx, p)
	if err != nil {
		return param.PlanExerciseResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.PlanExerciseResponse{
		ID: planExercise.ID,
		PlanID: planExercise.PlanID,
		ExerciseID: planExercise.ExerciseID,
		TargetSets: planExercise.TargetSets,
		TargetReps: planExercise.TargetReps,
		TargetWeight: planExercise.TargetWeight,
		SortOrder: planExercise.Order,
	}, nil

}