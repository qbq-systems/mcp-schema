package v20260728

type CallToolResultResponseClass struct {
	Description string                           `json:"description"`
	Properties  CallToolResultResponseProperties `json:"properties"`
	Required    []CallToolResultResponseRequired `json:"required"`
	Type        TypeElement                      `json:"type"`
}
