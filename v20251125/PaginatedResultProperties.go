package v20251125

type PaginatedResultProperties struct {
	Meta       GetTaskPayloadResult `json:"_meta"`
	NextCursor Cursor               `json:"nextCursor"`
}
