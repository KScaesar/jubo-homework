package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KScaesar/jubo-homework/backend/application"
	"github.com/KScaesar/jubo-homework/backend/domain"
)

func NewPatientHandler(svc application.PatientService) *PatientHandler {
	return &PatientHandler{svc: svc}
}

type PatientHandler struct {
	svc application.PatientService
}

func (h *PatientHandler) QueryPatientList(c *gin.Context) {
	var dto domain.QryPatientParam
	if !BindQueryStringOrPostFormRequest(c, &dto) {
		return
	}
	ctx := c.Request.Context()

	resp, err := h.svc.QueryPatientList(ctx, &dto)
	if err != nil {
		ReplyErrorResponse(c, err)
		return
	}

	ReplySuccessResponse(c, http.StatusOK, resp)
}
