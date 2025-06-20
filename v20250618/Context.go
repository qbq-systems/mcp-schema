package v20250618

type Context struct {
	Description string            `json:"description"`
	Properties  ContextProperties `json:"properties"`
	Type        AnnotationsType   `json:"type"`
}
