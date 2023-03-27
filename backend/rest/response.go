package rest

import (
	"github.com/gin-gonic/gin"

	"github.com/KScaesar/jubo-homework/backend/util/errors"
)

func ReplyErrorResponse(c *gin.Context, err error) {
	customError, _ := errors.ExtractCustomError(err)
	resp := &Response{
		Code:    customError.MyCode(),
		Message: err.Error(),
		Payload: struct{}{},
	}
	c.JSON(customError.HttpCode(), resp)
	c.Abort()
}

func ReplySuccessResponse(c *gin.Context, httpCode int, payload any) {
	if payload == nil {
		payload = struct{}{}
	}
	resp := &Response{
		Code:    0,
		Message: "ok",
		Payload: payload,
	}
	c.JSON(httpCode, resp)
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Payload any    `json:"payload"`
}
