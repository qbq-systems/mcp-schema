package v20251125

type Of struct {
	Description string       `json:"description"`
	Items       OneOfItems   `json:"items"`
	Type        AudienceType `json:"type"`
}
