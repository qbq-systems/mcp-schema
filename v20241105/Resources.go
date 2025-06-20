package v20241105

type Resources struct {
	Description string              `json:"description"`
	Properties  ResourcesProperties `json:"properties"`
	Type        AnnotationsType     `json:"type"`
}
