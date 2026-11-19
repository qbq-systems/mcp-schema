package v20260728

type ListRootsRequest struct {
	Description string                     `json:"description"`
	Properties  ListRootsRequestProperties `json:"properties"`
	Required    []CallToolRequestRequired  `json:"required"`
	Type        TypeElement                `json:"type"`
}
