package repository

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/workout-tracker.git/exercise/entity"
)

func (d *DB) CreateExercise(ctx context.Context, e entity.Exercise) error {
	var id uint
	query := "INSERT INTO exercises(name, muscle_group) VALUES($1, $2) RETURNING id"
	
	err := d.conn.Conn().QueryRow(ctx, query, e.Name, e.MuscleGroup).Scan(&id)
	if err != nil {
		return fmt.Errorf("can't insert into exercise table: %w", err)
	}

	e.ID = id

	return nil
}

func (d *DB) ListExercises(ctx context.Context) ([]entity.Exercise, error) {
	var exercise entity.Exercise
	var exercises []entity.Exercise
	query := "SELECT id, name, muscle_group FROM exercises"

	rows, err := d.conn.Conn().Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("can't list exercises %w", err)
	}

	for rows.Next() {
		rows.Scan(&exercise.ID, &exercise.Name, &exercise.MuscleGroup)
		exercises = append(exercises, exercise)
	}


	return exercises, nil

}
