package vdraft

type PaginatedRequest struct {
	Properties PaginatedRequestProperties `json:"properties"`
	Required   []CallToolRequestRequired  `json:"required"`
	Type       TypeElement                `json:"type"`
}
