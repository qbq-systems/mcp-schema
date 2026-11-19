package v20260728

type FluffyAllOf struct {
	Ref        *string              `json:"$ref,omitempty"`
	Properties *TentacledProperties `json:"properties,omitempty"`
	Required   []string             `json:"required,omitempty"`
	Type       *TypeElement         `json:"type,omitempty"`
}
