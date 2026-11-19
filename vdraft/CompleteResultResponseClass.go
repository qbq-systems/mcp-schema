package vdraft

type CompleteResultResponseClass struct {
	Description string                           `json:"description"`
	Properties  CompleteResultResponseProperties `json:"properties"`
	Required    []CallToolResultResponseRequired `json:"required"`
	Type        TypeElement                      `json:"type"`
}
