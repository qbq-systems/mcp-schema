package v20260728

type CallToolRequest struct {
	Description string                    `json:"description"`
	Properties  CallToolRequestProperties `json:"properties"`
	Required    []CallToolRequestRequired `json:"required"`
	Type        TypeElement               `json:"type"`
}
