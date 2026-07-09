package repository

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/workout-tracker.git/plan/entity"
)

func (d *DB) Add(ctx context.Context, p entity.WorkoutPlan) (entity.WorkoutPlan, error) {
	var id uint
	query := `INSERT INTO plans(user_id, name) VALUES($1, $2) RETURNING id`
	err := d.conn.Conn().QueryRow(ctx, query, p.UserID, p.Name).Scan(&id)
	if err != nil {
		return entity.WorkoutPlan{}, fmt.Errorf("can't insert into plans: %w", err)
	}

	p.ID = id
	
	return p, nil
}

func (d *DB) GetAll(ctx context.Context) ([]entity.WorkoutPlan, error) {
	var plan entity.WorkoutPlan
	var plans []entity.WorkoutPlan

	query := "SELECT id, user_id, name FROM plans"
	rows, err := d.conn.Conn().Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("can't get plans: %w", err)
	}

	for rows.Next() {
		rows.Scan(&plan.ID, &plan.UserID, &plan.Name)
		plans = append(plans, plan)
	}


	return plans, nil
}

func (d *DB) GetPlansOfUser(ctx context.Context, userID uint) ([]entity.WorkoutPlan, error) {
	var plan entity.WorkoutPlan
	var plans []entity.WorkoutPlan

	query := "SELECT id, user_id, name FROM plans WHERE user_id = $1"
	rows, err := d.conn.Conn().Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("can't get plans: %w", err)
	}

	for rows.Next() {
		rows.Scan(&plan.ID, &plan.UserID, &plan.Name)
		plans = append(plans, plan)
	}

	return plans, nil
}


func(d *DB) AddExercisePlan(ctx context.Context, p entity.PlanExercise) (entity.PlanExercise, error) {
	var id uint
	query := `INSERT INTO plan_exercises(plan_id, exercise_id, target_sets, target_reps, target_weight, sort_order)
				VALUES($1, $2, $3, $4, $5, $6) RETURNING id`
	err := d.conn.Conn().QueryRow(ctx, query, p.PlanID, p.ExerciseID, p.TargetSets, p.TargetReps, p.TargetWeight, p.Order).Scan(&id)
	if err != nil {
		return entity.PlanExercise{}, fmt.Errorf("can't insert into addexercise plan: %w", err)
	}

	p.ID = id

	return p, nil

}