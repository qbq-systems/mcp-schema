package vdraft

type SetLevelRequest struct {
	Description string                    `json:"description"`
	Properties  SetLevelRequestProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
