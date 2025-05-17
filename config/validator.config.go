package config

import "github.com/go-playground/validator/v10"

func ValidatorInt() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate
}
