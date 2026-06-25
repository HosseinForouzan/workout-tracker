package param

import "github.com/HosseinForouzan/workout-tracker.git/exercise/entity"

type ExerciseRequest struct {
	Name        string `json:"name"`
	MuscleGroup entity.MuscleGroup `json:"muscle_group"`
}