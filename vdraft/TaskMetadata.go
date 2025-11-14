package vdraft

type TaskMetadata struct {
	Description string                 `json:"description"`
	Properties  TaskMetadataProperties `json:"properties"`
	Type        AnnotationsType        `json:"type"`
}
