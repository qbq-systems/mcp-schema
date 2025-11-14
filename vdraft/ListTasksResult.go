package vdraft

type ListTasksResult struct {
	Description string                    `json:"description"`
	Properties  ListTasksResultProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
