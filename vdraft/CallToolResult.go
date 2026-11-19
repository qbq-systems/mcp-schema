package vdraft

type CallToolResult struct {
	Description string                   `json:"description"`
	Properties  CallToolResultProperties `json:"properties"`
	Required    []string                 `json:"required"`
	Type        TypeElement              `json:"type"`
}
