package vdraft

type SamplingMessage struct {
	Description string                    `json:"description"`
	Properties  SamplingMessageProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
