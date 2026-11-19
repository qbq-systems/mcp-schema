package v20260728

type PaginatedRequest struct {
	Properties PaginatedRequestProperties `json:"properties"`
	Required   []CallToolRequestRequired  `json:"required"`
	Type       TypeElement                `json:"type"`
}
