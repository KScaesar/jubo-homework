package errors

import (
	"errors"
	"fmt"
)

func ExtractCustomError(err error) (CustomError, bool) {
	var Err *CustomError
	if errors.As(err, &Err) {
		return *Err, true
	}
	return *ErrUnknown3rdParty, false
}

func Join3rdPartyWithMsg(myErr error, Err3rd error, msg string, args ...any) error {
	return fmt.Errorf("%v: %w: %w", fmt.Sprintf(msg, args...), Err3rd, myErr)
}

func Join3rdParty(myErr error, Err3rd error) error {
	return fmt.Errorf("%w: %w", Err3rd, myErr)
}

func WrapWithMessage(myErr error, msg string, args ...any) error {
	return fmt.Errorf("%v: %w", fmt.Sprintf(msg, args...), myErr)
}

func NewCustomError(title string, myCode int, httpCode int) *CustomError {
	return &CustomError{title: title, myCode: myCode, httpCode: httpCode}
}

type CustomError struct {
	title    string
	myCode   int
	httpCode int
}

func (c CustomError) Error() string {
	return c.title
}

func (c CustomError) MyCode() int {
	return c.myCode
}

func (c CustomError) HttpCode() int {
	return c.httpCode
}

func (c CustomError) CustomError() {}
