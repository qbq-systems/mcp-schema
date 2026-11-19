package vdraft

type Request struct {
	Description string                         `json:"description"`
	Properties  CreateMessageRequestProperties `json:"properties"`
	Required    []CallToolRequestRequired      `json:"required"`
	Type        TypeElement                    `json:"type"`
}
