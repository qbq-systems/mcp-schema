package vdraft

type Audience struct {
	Description string       `json:"description"`
	Items       EmptyResult  `json:"items"`
	Type        AudienceType `json:"type"`
}
