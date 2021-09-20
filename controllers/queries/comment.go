package queries

type CommentQuery struct {
	Pagination
}

type CommentPayload struct {
	Content string `json:"content" query:"content" validate:"required"`
}
