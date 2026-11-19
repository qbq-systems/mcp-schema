package v20260728

type Audience struct {
	Description string               `json:"description"`
	Items       AdditionalProperties `json:"items"`
	Type        AudienceType         `json:"type"`
}
