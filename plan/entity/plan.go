package entity

type WorkoutPlan struct {
	ID uint
	UserID uint
	Name string
	Exercises []PlanExercise
}

type PlanExercise struct {
    ExerciseID   uint
    TargetSets   uint
    TargetReps   uint
    TargetWeight float64
    Order        uint
}