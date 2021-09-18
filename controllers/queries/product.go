package queries

import "github.com/anti-lgbt/learning-be/types"

type ProductQuery struct {
	Type        string         `query:"type"`
	Name        string         `query:"name"`
	Special     bool           `query:"special"`
	OrderBy     string         `query:"order_by"`
	Ordering    types.Ordering `query:"ordering" default:"asc"`
	Discounting bool           `query:"discounting"`
	Pagination
}
