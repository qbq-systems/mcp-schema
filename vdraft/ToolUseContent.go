package vdraft

type ToolUseContent struct {
	Description string                   `json:"description"`
	Properties  ToolUseContentProperties `json:"properties"`
	Required    []string                 `json:"required"`
	Type        AnnotationsType          `json:"type"`
}
