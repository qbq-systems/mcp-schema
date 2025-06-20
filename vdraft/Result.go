package vdraft

type Result struct {
	AdditionalProperties AdditionalPropertiesClass `json:"additionalProperties"`
	Properties           MagentaProperties         `json:"properties"`
	Type                 AnnotationsType           `json:"type"`
}
