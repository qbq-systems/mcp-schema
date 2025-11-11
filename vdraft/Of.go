package vdraft

type Of struct {
	Description string       `json:"description"`
	Items       OneOfItems   `json:"items"`
	Type        AudienceType `json:"type"`
}
