package vdraft

type ListToolsResult struct {
	Description string                    `json:"description"`
	Properties  ListToolsResultProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
