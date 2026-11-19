package vdraft

type ToolChoice struct {
	Description string               `json:"description"`
	Properties  ToolChoiceProperties `json:"properties"`
	Type        TypeElement          `json:"type"`
}
