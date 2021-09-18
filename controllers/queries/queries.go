package queries

type Pagination struct {
	Limit int `query:"limit" validate:"uint|default:100"`
	Page  int `query:"page" validate:"uint|default:1"`
}

type Period struct {
	TimeFrom int64 `query:"time_from" validate:"uint"`
	TimeTo   int64 `query:"time_to" validate:"uint"`
}
