package vdraft

type PaginatedRequestParamsProperties struct {
	Meta   AdditionalProperties `json:"_meta"`
	Cursor Cursor               `json:"cursor"`
}
