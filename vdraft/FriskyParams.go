package vdraft

type FriskyParams struct {
	Properties Properties1     `json:"properties"`
	Required   []FormatElement `json:"required"`
	Type       AnnotationsType `json:"type"`
}
