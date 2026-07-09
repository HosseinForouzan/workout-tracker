package param

import "github.com/HosseinForouzan/workout-tracker.git/plan/entity"

type PlanAddRequest struct {
	UserID uint `json:"user_id"`
	Name   string `json:"name"`
}

type PlanAddResponse struct {
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`
	Name   string `json:"name"`
}

type PlansGetResponse struct {
	WorkoutPlans []entity.WorkoutPlan
}

type PlansOfUserGetResponse struct {
	WorkoutPlans []entity.WorkoutPlan
}

type PlanExerciseRequest struct {
	PlanID uint `json:"plan_id"`
	ExerciseID uint `json:"exercise_id"`
	TargetSets uint `json:"target_sets"`
	TargetReps uint `json:"target_reps"`
	TargetWeight float64 `json:"target_weight"`
	SortOrder uint `json:"sort_order"`
}

type PlanExerciseResponse struct {
	ID uint `json:"id"`
	PlanID uint `json:"plan_id"`
	ExerciseID uint `json:"exercise_id"`
	TargetSets uint `json:"target_sets"`
	TargetReps uint `json:"target_reps"`
	TargetWeight float64 `json:"target_weight"`
	SortOrder uint `json:"sort_order"`
}