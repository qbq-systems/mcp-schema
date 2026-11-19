package v20260728

type ContentAnyOf struct {
	Ref   *string               `json:"$ref,omitempty"`
	Items *AdditionalProperties `json:"items,omitempty"`
	Type  *AudienceType         `json:"type,omitempty"`
}
