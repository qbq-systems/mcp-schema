package v20241105

type Annotated struct {
	Description string              `json:"description"`
	Properties  AnnotatedProperties `json:"properties"`
	Type        AnnotationsType     `json:"type"`
}
