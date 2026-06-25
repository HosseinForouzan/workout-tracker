package entity

import "time"

type WorkoutSession struct {
	ID        uint
	UserID    uint
	PlanID    uint
	Status    string
	StartedAt time.Time
	CompletedAt time.Time
	Sets        []WorkoutSet
}

type WorkoutSet struct {
	ExerciseID uint
	SetNumber  uint
	Reps       uint
	Weight     float64
}