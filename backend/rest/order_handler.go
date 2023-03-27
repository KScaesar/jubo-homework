package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KScaesar/jubo-homework/backend/application"
	"github.com/KScaesar/jubo-homework/backend/domain"
	"github.com/KScaesar/jubo-homework/backend/util/database"
)

func NewOrderHandler(svc application.OrderService, txFactory database.TransactionFactory) *OrderHandler {
	return &OrderHandler{svc: svc, txFactory: txFactory}
}

type OrderHandler struct {
	svc       application.OrderService
	txFactory database.TransactionFactory
}

func (h *OrderHandler) QueryOrderList(c *gin.Context) {
	var dto domain.DtoQryOrderParam
	if !BindQueryStringOrPostFormRequest(c, &dto) {
		return
	}
	ctx := c.Request.Context()

	resp, err := h.svc.QueryOrderList(ctx, &dto)
	if err != nil {
		ReplyErrorResponse(c, err)
		return
	}

	ReplySuccessResponse(c, http.StatusOK, resp)
}

func (h *OrderHandler) QueryOrderById(c *gin.Context) {
	ctx := c.Request.Context()
	orderId, _ := c.Params.Get("order_id")

	resp, err := h.svc.QueryOrderById(ctx, orderId)
	if err != nil {
		ReplyErrorResponse(c, err)
		return
	}

	ReplySuccessResponse(c, http.StatusOK, resp)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var dto domain.DtoCreateOrder
	if !BindJsonRequest(c, &dto) {
		return
	}
	ctx := c.Request.Context()

	if err := h.txFactory.CreateTransaction(ctx).AutoCommit(
		func(ctx context.Context) error {
			_, err := h.svc.CreateOrder(ctx, &dto)
			return err
		},
	); err != nil {
		ReplyErrorResponse(c, err)
		return
	}

	ReplySuccessResponse(c, http.StatusOK, nil)
}

func (h *OrderHandler) UpdateOrderInfo(c *gin.Context) {
	var dto domain.DtoUpdateOrderInfo
	if !BindJsonRequest(c, &dto) {
		return
	}
	ctx := c.Request.Context()
	orderId, _ := c.Params.Get("order_id")

	if err := h.txFactory.CreateTransaction(ctx).AutoCommit(
		func(ctx context.Context) error {
			return h.svc.UpdateOrderInfo(ctx, orderId, &dto)
		},
	); err != nil {
		ReplyErrorResponse(c, err)
		return
	}

	ReplySuccessResponse(c, http.StatusOK, nil)
}
