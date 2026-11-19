package vdraft

type TentacledAllOf struct {
	Ref        *string            `json:"$ref,omitempty"`
	Properties *CunningProperties `json:"properties,omitempty"`
	Required   []string           `json:"required,omitempty"`
	Type       *TypeElement       `json:"type,omitempty"`
}
