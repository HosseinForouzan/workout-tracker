package repository

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/workout-tracker.git/user/entity"
)

func (d *DB) Register(ctx context.Context, user entity.User) (entity.User, error) {
	var id uint
	query := `INSERT INTO users(name, email, password) 
			VALUES($1, $2, $3) RETURNING id`
	err := d.conn.Conn().QueryRow(ctx, query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't insert into register table: %w", err)
	}

	user.ID = id

	return user, nil
}