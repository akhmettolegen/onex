package handlers

import (
	"github.com/akhmettolegen/texert/pkg/helpers"
	"github.com/akhmettolegen/texert/pkg/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// GetProducts godoc
// @Tags Products
// @Summary Retrieve Products list from database
// @ID get-products
// @Security ApiKeyAuth
// @Accept json
// @Param me	query string false "Get user's products if true"
// @Param status	query string false "statuses list by commas (READY, PENDING)"
// @Param page	query int false "Page number" default(1)
// @Param size	query int false "Page size" default(15)
// @Produce json
// @Success 200 {object} models.ProductsListResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /products [get]
func (h *Handler) GetProducts(ctx *gin.Context) {
	ui := ctx.MustGet(models.UserInfoKey).(*models.UserInfo)
	var query helpers.RequestQuery
	err := ctx.Bind(&query)
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	page, size := helpers.ParsePagination(query)
	statusFilters := helpers.GetStatusFiltersFromQuery(ctx)

	response, err := h.Manager.GetProducts(ui, page, size, statusFilters)
	if err != nil {
		ctx.JSON(500, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, response)
}

// GetProductByID godoc
// @Tags Products
// @Summary Retrieve product by id from database
// @ID get-product-by-id
// @Security ApiKeyAuth
// @Accept json
// @Param id path string true "product ID"
// @Produce json
// @Success 200 {object} models.ProductByIDResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /products/{id} [get]
func (h *Handler) GetProductByID(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}
	response, err := h.Manager.GetProductByID(id)
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, response)
}

// CreateProduct godoc
// @Tags Products
// @Summary Creates product
// @ID create-product
// @Security ApiKeyAuth
// @Accept json
// @Param product body models.ProductCreateRequest true "product body"
// @Produce json
// @Success 200 {object} models.ProductByIDResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /products [post]
func (h *Handler) CreateProduct(ctx *gin.Context) {
	ui := ctx.MustGet(models.UserInfoKey).(*models.UserInfo)

	var productReq models.ProductCreateRequest
	err := ctx.Bind(&productReq)
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	response, err := h.Manager.CreateProduct(ui, productReq)
	if err != nil {
		ctx.JSON(500, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(201, response)
}

// UpdateProductByID godoc
// @Tags Products
// @Summary Updates specific product by id
// @ID update-product-by-id
// @Security ApiKeyAuth
// @Accept json
// @Param id path string true "product ID"
// @Param product body models.ProductUpdateRequest true "product body"
// @Produce json
// @Success 200 {object} models.ProductByIDResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /products/{id} [put]
func (h *Handler) UpdateProductByID(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	var productData models.ProductUpdateRequest
	if err := ctx.BindJSON(&productData); err != nil {
		ctx.JSON(400, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	productData.ID = id
	response, err := h.Manager.UpdateProduct(productData)
	if err != nil {
		ctx.JSON(500, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, response)
}

// DeleteProduct godoc
// @Tags Products
// @Summary Deletes specific product by id
// @ID delete-product-by-id
// @Security ApiKeyAuth
// @Accept json
// @Param id path string true "product ID"
// @Produce json
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /products/{id} [delete]
func (h *Handler) DeleteProduct(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	err = h.Manager.DeleteProduct(id)
	if err != nil {
		ctx.JSON(500, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, models.BaseResponse{
		Message: "product deleted",
	})
}