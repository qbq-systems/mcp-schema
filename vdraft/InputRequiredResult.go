package vdraft

type InputRequiredResult struct {
	Description string                        `json:"description"`
	Properties  InputRequiredResultProperties `json:"properties"`
	Required    []string                      `json:"required"`
	Type        TypeElement                   `json:"type"`
}
