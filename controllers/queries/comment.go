package queries

type CommentQuery struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
