package validator

import (
	"encoding/json"
	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4"

	"github.com/HosseinForouzan/workout-tracker.git/user/param"
)

func (v Validator) ValidateRegisterRequest(req param.RegisterReqeust) (string, error) {
	err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.Email, validation.Required,
			 validation.Match(regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`))),
		validation.Field(&req.Password, validation.Required, validation.Match(regexp.MustCompile(`^.{4,}$`))),
	)

	if err != nil {
		
		b, _ := json.Marshal(err)

		return string(b), err
	}

	return "", nil

}