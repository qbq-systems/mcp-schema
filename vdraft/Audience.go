package vdraft

type Audience struct {
	Description string               `json:"description"`
	Items       AdditionalProperties `json:"items"`
	Type        AudienceType         `json:"type"`
}
