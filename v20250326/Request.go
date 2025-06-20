package v20250326

type Request struct {
	Description string                     `json:"description"`
	Properties  ListRootsRequestProperties `json:"properties"`
	Required    []string                   `json:"required"`
	Type        AnnotationsType            `json:"type"`
}
