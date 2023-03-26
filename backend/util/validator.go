package util

import "github.com/go-playground/validator/v10"

var (
	Validator = validator.New()
)

func init() {
	// https://stackoverflow.com/questions/71278746/validate-enum-in-golang-using-gin-framework
	Validator.RegisterValidation(NewRegisterValidation[SortKind]("sort"))
}

func NewRegisterValidation[T interface{ IsValid() bool }](tagKeyName string) (string, func(fl validator.FieldLevel) bool) {
	playground_validator_signature := func(fl validator.FieldLevel) bool {
		value := fl.Field().Interface().(T)
		return value.IsValid()
	}
	return tagKeyName, playground_validator_signature
}
