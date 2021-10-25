package component

type Pagination struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Total int64       `json:"total"`
	Items interface{} `json:"items"`
}

// CalOffset 计算offset /**
func CalOffset(page int, limit int) (offset int) {
	if page > 0 {
		offset = (page - 1) * limit
	}

	return offset
}

// PageStructure 构建分页数据结构 **/
func PageStructure(page int, limit int, total int64, items interface{}) *Pagination {
	return &Pagination{
		Page:  page,
		Limit: limit,
		Total: total,
		Items: items,
	}
}
