package vdraft

type PaginatedResultProperties struct {
	Meta       AdditionalProperties `json:"_meta"`
	NextCursor Cursor               `json:"nextCursor"`
	ResultType Cursor               `json:"resultType"`
}
