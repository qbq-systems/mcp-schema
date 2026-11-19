package v20260728

type ToolChoice struct {
	Description string               `json:"description"`
	Properties  ToolChoiceProperties `json:"properties"`
	Type        TypeElement          `json:"type"`
}
