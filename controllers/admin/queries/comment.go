package queries

import "github.com/anti-lgbt/learning-be/controllers/queries"

type CommentQuery struct {
	UserID    uint64 `query:"page" validate:"uint"`
	ProductID uint64 `query:"page" validate:"uint"`
	queries.Order
	queries.Pagination
	queries.Period
}
