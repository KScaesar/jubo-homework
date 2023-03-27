package util

import "github.com/go-playground/validator/v10"

var (
	Validator = validator.New()
)

func init() {
	// https://stackoverflow.com/questions/71278746/validate-enum-in-golang-using-gin-framework
	Validator.RegisterValidation(NewRegisterValidation[SortKind]("sort"))
}

type IValidator interface {
	Validate() error
}

func NewRegisterValidation[T IValidator](tagKeyName string) (string, validator.Func) {
	return tagKeyName, func(fl validator.FieldLevel) bool {
		obj := fl.Field().Interface().(T)
		err := obj.Validate()
		if err != nil {
			return false
		}
		return true
	}
}
