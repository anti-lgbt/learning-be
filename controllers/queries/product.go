package queries

type ProductQuery struct {
	Type string `query:"type"`
	Name string `query:"name"`
	Pagination
}
