package vdraft

type AdditionalPropertiesAnyOf struct {
	Items *Default   `json:"items,omitempty"`
	Type  *TypeUnion `json:"type"`
}
