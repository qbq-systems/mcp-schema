package v20260728

type Request struct {
	Description string                         `json:"description"`
	Properties  CreateMessageRequestProperties `json:"properties"`
	Required    []CallToolRequestRequired      `json:"required"`
	Type        TypeElement                    `json:"type"`
}
