package vdraft

type PaginatedResult struct {
	Properties PaginatedResultProperties `json:"properties"`
	Required   []string                  `json:"required"`
	Type       TypeElement               `json:"type"`
}
