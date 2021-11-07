package helpers

import (
	"github.com/akhmettolegen/onex/pkg/models"
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

func GetStatusFiltersFromQuery(ctx *gin.Context) []models.OrderStatus {
	result := []models.OrderStatus{models.OrderStatusRecommended, models.OrderStatusPending, models.OrderStatusReady}

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

func spaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}