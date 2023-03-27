package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/KScaesar/jubo-homework/backend/util"
	"github.com/KScaesar/jubo-homework/backend/util/errors"
)

// https://gin-gonic.com/docs/examples/binding-and-validation/

func BindJsonRequest(c *gin.Context, obj any) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		Err := errors.Join3rdPartyWithMsg(errors.ErrInvalidParams, err, "bind json payload")
		ReplyErrorResponse(c, Err)
		return false
	}
	return ValidateDtoRequest(c, obj)
}

func BindQueryStringOrPostFormRequest(c *gin.Context, obj any) bool {
	if err := c.ShouldBind(obj); err != nil {
		Err := errors.Join3rdPartyWithMsg(errors.ErrInvalidParams, err, "bind query string or form payload")
		ReplyErrorResponse(c, Err)
		return false
	}
	return ValidateDtoRequest(c, obj)
}

func ValidateDtoRequest(c *gin.Context, obj any) bool {
	var target *validator.InvalidValidationError
	err := util.Validator.StructCtx(c.Request.Context(), obj)
	if err != nil {
		if errors.As(err, &target) {
			return true
		}

		Err := errors.Join3rdParty(errors.ErrInvalidParams, err)
		ReplyErrorResponse(c, Err)
		return false
	}
	return true
}
