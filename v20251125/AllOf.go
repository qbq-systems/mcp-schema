package v20251125

type AllOf struct {
	Ref        *string          `json:"$ref,omitempty"`
	Properties *AllOfProperties `json:"properties,omitempty"`
	Required   []string         `json:"required,omitempty"`
	Type       *AnnotationsType `json:"type,omitempty"`
}
