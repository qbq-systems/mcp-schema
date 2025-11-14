package vdraft

type PaginatedRequestParamsProperties struct {
	Meta   GetTaskPayloadResult `json:"_meta"`
	Cursor Cursor               `json:"cursor"`
	Task   AnnotationsElement   `json:"task"`
}
