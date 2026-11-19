package v20260728

type PaginatedResultProperties struct {
	Meta       AdditionalProperties `json:"_meta"`
	NextCursor Cursor               `json:"nextCursor"`
	ResultType Cursor               `json:"resultType"`
}
