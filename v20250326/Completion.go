package v20250326

type Completion struct {
	Properties CompletionProperties `json:"properties"`
	Required   []string             `json:"required"`
	Type       AnnotationsType      `json:"type"`
}
