package queries

type ProductQuery struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}
