package v20241105

type ResourceContents struct {
	Description string                     `json:"description"`
	Properties  ResourceContentsProperties `json:"properties"`
	Required    []Required                 `json:"required"`
	Type        AnnotationsType            `json:"type"`
}
