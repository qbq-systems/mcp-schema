package vdraft

type JSONValueAnyOf struct {
	Ref   *string               `json:"$ref,omitempty"`
	Items *AdditionalProperties `json:"items,omitempty"`
	Type  *TypeUnion            `json:"type,omitempty"`
}
