package vdraft

type Values struct {
	Description string       `json:"description"`
	Items       Default      `json:"items"`
	Type        AudienceType `json:"type"`
}
