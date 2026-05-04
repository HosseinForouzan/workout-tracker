package repository

import (
	"context"
	"database/sql"
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

func (d *DB) DoesUserExistByEmail(ctx context.Context, email string) (bool, error) {
	var user entity.User

	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	err := d.conn.Conn().QueryRow(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("can't get existence of user: %w", err)
	}

	return true, nil
}

func (d *DB) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User

	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	err := d.conn.Conn().QueryRow(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't get user by email: %w", err)
	}

	return user, nil
}