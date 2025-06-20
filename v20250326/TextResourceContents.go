package v20250326

type TextResourceContents struct {
	Properties TextResourceContentsProperties `json:"properties"`
	Required   []string                       `json:"required"`
	Type       AnnotationsType                `json:"type"`
}
