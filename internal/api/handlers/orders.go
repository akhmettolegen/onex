package handlers

import (
	"github.com/akhmettolegen/onex/pkg/helpers"
	"github.com/akhmettolegen/onex/pkg/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// GetOrders godoc
// @Tags Orders
// @Summary Retrieve orders list from database
// @ID get-orders
// @Security ApiKeyAuth
// @Accept json
// @Param me	query string false "Get user's orders if true"
// @Param status	query string false "statuses list by commas (READY, PENDING)"
// @Param page	query int false "Page number" default(1)
// @Param size	query int false "Page size" default(15)
// @Produce json
// @Success 200 {object} models.OrdersListResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /orders [get]
func (h *Handler) GetOrders(ctx *gin.Context) {
	ui := ctx.MustGet(models.UserInfoKey).(*models.UserInfo)
	var query helpers.RequestQuery
	err := ctx.Bind(&query)
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Code:    0,
			Message: err.Error(),
		})
		return
	}

	page, size := helpers.ParsePagination(query)
	statusFilters := helpers.GetStatusFiltersFromQuery(ctx)
	me := ctx.Query("me")

	response, err := h.Manager.GetOrders(ui, page, size, me, statusFilters)
	if err != nil {
		ctx.JSON(500, models.BaseResponse{
			Code:    0,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, response)
}

// GetOrderByID godoc
// @Tags Orders
// @Summary Retrieve order by id from database
// @ID get-order-by-id
// @Security ApiKeyAuth
// @Accept json
// @Param id path string true "Order ID"
// @Produce json
// @Success 200 {object} models.OrderByIDResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /orders/{id} [get]
func (h *Handler) GetOrderByID(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Code:    0,
			Message: err.Error(),
		})
		return
	}
	response, err := h.Manager.GetOrderByID(id)
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Code:    0,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, response)
}

// CreateOrder godoc
// @Tags Orders
// @Summary Creates order
// @ID create-order
// @Security ApiKeyAuth
// @Accept json
// @Param order body models.OrderCreateRequest true "Order body"
// @Produce json
// @Success 200 {object} models.OrderByIDResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /orders [post]
func (h *Handler) CreateOrder(ctx *gin.Context) {
	ui := ctx.MustGet(models.UserInfoKey).(*models.UserInfo)

	var orderReq models.OrderCreateRequest
	err := ctx.Bind(&orderReq)
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Code:    0,
			Message: err.Error(),
		})
		return
	}

	response, err := h.Manager.CreateOrder(ui, orderReq)
	if err != nil {
		ctx.JSON(500, models.BaseResponse{
			Code:    0,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(201, response)
}

// UpdateOrderByID godoc
// @Tags Orders
// @Summary Updates specific order by id
// @ID update-order-by-id
// @Security ApiKeyAuth
// @Accept json
// @Param id path string true "Order ID"
// @Param order body models.OrderUpdateRequest true "Order body"
// @Produce json
// @Success 200 {object} models.OrderByIDResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /orders/{id} [put]
func (h *Handler) UpdateOrderByID(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Code:    0,
			Message: err.Error(),
		})
		return
	}

	var orderData models.OrderUpdateRequest
	if err := ctx.BindJSON(&orderData); err != nil {
		ctx.JSON(400, models.BaseResponse{
			Code:    0,
			Message: err.Error(),
		})
		return
	}

	orderData.ID = id
	response, err := h.Manager.UpdateOrder(orderData)
	if err != nil {
		ctx.JSON(500, models.BaseResponse{
			Code:    0,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, response)
}

// DeleteOrder godoc
// @Tags Orders
// @Summary Deletes specific order by id
// @ID delete-order-by-id
// @Security ApiKeyAuth
// @Accept json
// @Param id path string true "Order ID"
// @Produce json
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /orders/{id} [delete]
func (h *Handler) DeleteOrder(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Code:    0,
			Message: err.Error(),
		})
		return
	}

	err = h.Manager.DeleteOrder(id)
	if err != nil {
		ctx.JSON(500, models.BaseResponse{
			Code:    0,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, models.BaseResponse{
		Code:    0,
		Message: "order deleted",
	})
}