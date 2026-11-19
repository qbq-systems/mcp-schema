package v20251125

type ToolChoice struct {
	Description string               `json:"description"`
	Properties  ToolChoiceProperties `json:"properties"`
	Type        AnnotationsType      `json:"type"`
}
