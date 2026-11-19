package vdraft

type PurpleAllOf struct {
	Ref        *string           `json:"$ref,omitempty"`
	Properties *FluffyProperties `json:"properties,omitempty"`
	Required   []ErrorRequired   `json:"required,omitempty"`
	Type       *TypeElement      `json:"type,omitempty"`
}
