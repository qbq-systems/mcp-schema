package v20241105

type CallToolRequest struct {
	Description string                    `json:"description"`
	Properties  CallToolRequestProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
