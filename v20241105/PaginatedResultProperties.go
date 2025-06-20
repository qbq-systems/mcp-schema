package v20241105

type PaginatedResultProperties struct {
	Meta       PurpleMeta `json:"_meta"`
	NextCursor Cursor     `json:"nextCursor"`
}
