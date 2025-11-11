package vdraft

type AnyOf struct {
	Items *Default   `json:"items,omitempty"`
	Type  *TypeUnion `json:"type"`
}
