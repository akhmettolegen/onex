package handlers

import (
	"github.com/akhmettolegen/texert/pkg/helpers"
	"github.com/akhmettolegen/texert/pkg/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// GetCategories godoc
// @Tags Categories
// @Summary Retrieve Categories list from database
// @ID get-categories
// @Security ApiKeyAuth
// @Accept json
// @Param me	query string false "Get user's categories if true"
// @Param status	query string false "statuses list by commas (READY, PENDING)"
// @Param page	query int false "Page number" default(1)
// @Param size	query int false "Page size" default(15)
// @Produce json
// @Success 200 {object} models.CategoriesListResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /categories [get]
func (h *Handler) GetCategories(ctx *gin.Context) {
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

	response, err := h.Manager.GetCategories(ui, page, size)
	if err != nil {
		ctx.JSON(500, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, response)
}

// GetCategoryByID godoc
// @Tags Categories
// @Summary Retrieve category by id from database
// @ID get-category-by-id
// @Security ApiKeyAuth
// @Accept json
// @Param id path string true "category ID"
// @Produce json
// @Success 200 {object} models.CategoryByIDResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /categories/{id} [get]
func (h *Handler) GetCategoryByID(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}
	response, err := h.Manager.GetCategoryByID(id)
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, response)
}

// CreateCategory godoc
// @Tags Categories
// @Summary Creates category
// @ID create-category
// @Security ApiKeyAuth
// @Accept json
// @Param category body models.CategoryCreateRequest true "category body"
// @Produce json
// @Success 200 {object} models.CategoryByIDResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /categories [post]
func (h *Handler) CreateCategory(ctx *gin.Context) {
	ui := ctx.MustGet(models.UserInfoKey).(*models.UserInfo)

	var categoryReq models.CategoryCreateRequest
	err := ctx.Bind(&categoryReq)
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	response, err := h.Manager.CreateCategory(ui, categoryReq)
	if err != nil {
		ctx.JSON(500, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(201, response)
}

// UpdateCategoryByID godoc
// @Tags Categories
// @Summary Updates specific category by id
// @ID update-category-by-id
// @Security ApiKeyAuth
// @Accept json
// @Param id path string true "category ID"
// @Param category body models.CategoryUpdateRequest true "category body"
// @Produce json
// @Success 200 {object} models.CategoryByIDResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /categories/{id} [put]
func (h *Handler) UpdateCategoryByID(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	var categoryData models.CategoryUpdateRequest
	if err := ctx.BindJSON(&categoryData); err != nil {
		ctx.JSON(400, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	categoryData.ID = id
	response, err := h.Manager.UpdateCategory(categoryData)
	if err != nil {
		ctx.JSON(500, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, response)
}

// DeleteCategory godoc
// @Tags Categories
// @Summary Deletes specific category by id
// @ID delete-category-by-id
// @Security ApiKeyAuth
// @Accept json
// @Param id path string true "category ID"
// @Produce json
// @Success 200 {object} models.BaseResponse
// @Failure 400 {object} models.BaseResponse
// @Failure 500 {string} models.BaseResponse
// @Router /categories/{id} [delete]
func (h *Handler) DeleteCategory(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	err = h.Manager.DeleteCategory(id)
	if err != nil {
		ctx.JSON(500, models.BaseResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, models.BaseResponse{
		Message: "category deleted",
	})
}