package vdraft

type RelatedTaskMetadata struct {
	Properties  RelatedTaskMetadataProperties `json:"properties"`
	Required    []string                      `json:"required"`
	Type        AnnotationsType               `json:"type"`
	Description *string                       `json:"description,omitempty"`
}
