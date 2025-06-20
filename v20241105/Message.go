package v20241105

type Message struct {
	Description string                  `json:"description"`
	Properties  PromptMessageProperties `json:"properties"`
	Required    []string                `json:"required"`
	Type        AnnotationsType         `json:"type"`
}
