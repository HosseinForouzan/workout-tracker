package validator

import "context"

type Repository interface {
	DoesUserExistByEmail(ctx context.Context, email string) (bool, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{
		repo: repo,
	}
}