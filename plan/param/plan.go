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