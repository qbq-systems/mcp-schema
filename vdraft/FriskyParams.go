package vdraft

type FriskyParams struct {
	Properties Properties1     `json:"properties"`
	Required   []Required      `json:"required"`
	Type       AnnotationsType `json:"type"`
}
