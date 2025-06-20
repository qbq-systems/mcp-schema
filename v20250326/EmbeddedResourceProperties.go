package v20250326

type EmbeddedResourceProperties struct {
	Annotations AnnotationsClass   `json:"annotations"`
	Resource    ClientNotification `json:"resource"`
	Type        TypeClass          `json:"type"`
}
