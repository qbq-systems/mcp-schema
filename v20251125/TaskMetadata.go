package v20251125

type TaskMetadata struct {
	Description string                 `json:"description"`
	Properties  TaskMetadataProperties `json:"properties"`
	Type        AnnotationsType        `json:"type"`
}
