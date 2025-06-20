package v20250618

type CreateMessageResult struct {
	Description string                        `json:"description"`
	Properties  CreateMessageResultProperties `json:"properties"`
	Required    []string                      `json:"required"`
	Type        AnnotationsType               `json:"type"`
}
