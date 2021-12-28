package helpers

import (
	"github.com/akhmettolegen/texert/pkg/models"
	"github.com/gin-gonic/gin"
	"strings"
	"unicode"
)

type RequestQuery struct {
	Page 	*int		`form:"page" default:"1" json:"page"`
	Size 	*int 		`form:"size" default:"15" json:"size"`
}

func ParsePagination(query RequestQuery) (int, int) {
	page := 1
	if query.Page != nil {
		page = *query.Page
	}

	size := 15
	if query.Size != nil {
		size = *query.Size
	}
	return page, size
}

func GetStatusFiltersFromQueryOrder(ctx *gin.Context) []models.OrderStatus {
	result := []models.OrderStatus{models.OrderStatusNotAnalyzed, models.OrderStatusAnalyzed}

	qStatus := ctx.Query("status")
	if len(qStatus) > 0 {
		result = []models.OrderStatus{}
		separated := strings.Split(qStatus, ",")

		for _, s := range separated {
			trimmedStr := spaceMap(s)

			result = append(result, models.OrderStatus(trimmedStr))
		}
	}

	return result
}

func GetStatusFiltersFromQuery(ctx *gin.Context) []models.ProductStatus {
	result := []models.ProductStatus{models.ProductStatusActive, models.ProductStatusInactive}

	qStatus := ctx.Query("status")
	if len(qStatus) > 0 {
		result = []models.ProductStatus{}
		separated := strings.Split(qStatus, ",")

		for _, s := range separated {
			trimmedStr := spaceMap(s)

			result = append(result, models.ProductStatus(trimmedStr))
		}
	}

	return result
}

func spaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}