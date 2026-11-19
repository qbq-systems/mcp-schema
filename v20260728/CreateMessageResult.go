package v20260728

type CreateMessageResult struct {
	Description string                        `json:"description"`
	Properties  CreateMessageResultProperties `json:"properties"`
	Required    []string                      `json:"required"`
	Type        TypeElement                   `json:"type"`
}
