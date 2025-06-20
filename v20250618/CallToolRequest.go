package v20250618

type CallToolRequest struct {
	Description string                    `json:"description"`
	Properties  CallToolRequestProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
