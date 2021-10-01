package helpers

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
