package handlers

import (
	"github.com/akhmettolegen/onex/pkg/helpers"
	"github.com/akhmettolegen/onex/pkg/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func (h *Handler) GetOrders(ctx *gin.Context) {
	ui := ctx.MustGet(models.UserInfoKey).(*models.UserInfo)
	var query helpers.RequestQuery
	err := ctx.Bind(&query)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	page, size := helpers.ParsePagination(query)
	statusFilters := helpers.GetStatusFiltersFromQuery(ctx)
	me := ctx.Query("me")

	response, err := h.Manager.GetOrders(ui, page, size, me, statusFilters)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, response)
}

func (h *Handler) GetOrderByID(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	response, err := h.Manager.GetOrderByID(id)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, response)
}

func (h *Handler) CreateOrder(ctx *gin.Context) {
	ui := ctx.MustGet(models.UserInfoKey).(*models.UserInfo)
	// Parse request body
	var orderReq models.OrderCreateRequest
	err := ctx.Bind(&orderReq)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	response, err := h.Manager.CreateOrder(ui, orderReq)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(201, response)
}

func (h *Handler) UpdateOrderByID(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	var orderData models.OrderUpdateRequest
	if err := ctx.BindJSON(&orderData); err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	orderData.ID = id
	response, err := h.Manager.UpdateOrder(orderData)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, response)
}

func (h *Handler) DeleteOrder(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err = h.Manager.DeleteOrder(id)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "order deleted"})
}