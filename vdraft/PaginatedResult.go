package vdraft

type PaginatedResult struct {
	Properties PaginatedResultProperties `json:"properties"`
	Type       AnnotationsType           `json:"type"`
}
