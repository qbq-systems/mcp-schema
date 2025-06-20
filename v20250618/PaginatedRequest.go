package v20250618

type PaginatedRequest struct {
	Properties PaginatedRequestProperties `json:"properties"`
	Required   []string                   `json:"required"`
	Type       AnnotationsType            `json:"type"`
}
