package pagination

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewPagination(page int, limit int) *Pagination {
	return &Pagination{
		Page:  page,
		Limit: limit,
	}
}
