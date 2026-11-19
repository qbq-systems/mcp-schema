package v20251125

type ContentAnyOf struct {
	Ref   *string       `json:"$ref,omitempty"`
	Items *EmptyResult  `json:"items,omitempty"`
	Type  *AudienceType `json:"type,omitempty"`
}
