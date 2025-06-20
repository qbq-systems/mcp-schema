package v20241105

type Request struct {
	Description string                     `json:"description"`
	Properties  ListRootsRequestProperties `json:"properties"`
	Required    []string                   `json:"required"`
	Type        AnnotationsType            `json:"type"`
}
