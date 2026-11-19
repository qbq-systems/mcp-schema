package v20260728

type ToolUseContent struct {
	Description string                   `json:"description"`
	Properties  ToolUseContentProperties `json:"properties"`
	Required    []string                 `json:"required"`
	Type        TypeElement              `json:"type"`
}
