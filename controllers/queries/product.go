package queries

type ProductQuery struct {
	Type        string `query:"type"`
	Name        string `query:"name"`
	Special     bool   `query:"special"`
	Discounting bool   `query:"discounting"`
	Order
	Pagination
}
