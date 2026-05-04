package validator

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/HosseinForouzan/workout-tracker.git/user/param"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateLoginRequest(req param.LoginRequest) (string, error) {
err := validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required,
			 validation.Match(regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)),
			validation.By(v.doesEmailExist)),
		validation.Field(&req.Password, validation.Required, validation.Match(regexp.MustCompile(`^.{4,}$`))),
	)

	if err != nil {
		b, _ := json.Marshal(err)
		return string(b), err
	}

	return "", nil
			
}

func (v Validator) doesEmailExist(value any) error {
	email := value.(string)
	_, err := v.repo.DoesUserExistByEmail(context.Background(),email)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	return nil
}