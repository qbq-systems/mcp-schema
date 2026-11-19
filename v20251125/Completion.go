package v20251125

type Completion struct {
	Properties CompletionProperties `json:"properties"`
	Required   []string             `json:"required"`
	Type       AnnotationsType      `json:"type"`
}
