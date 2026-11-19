package v20251125

type Audience struct {
	Description string       `json:"description"`
	Items       EmptyResult  `json:"items"`
	Type        AudienceType `json:"type"`
}
