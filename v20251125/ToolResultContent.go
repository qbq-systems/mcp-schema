package v20251125

type ToolResultContent struct {
	Description string                      `json:"description"`
	Properties  ToolResultContentProperties `json:"properties"`
	Required    []string                    `json:"required"`
	Type        AnnotationsType             `json:"type"`
}
