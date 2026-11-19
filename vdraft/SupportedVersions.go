package vdraft

type SupportedVersions struct {
	Description string       `json:"description"`
	Items       Default      `json:"items"`
	MaxItems    *int64       `json:"maxItems,omitempty"`
	Type        AudienceType `json:"type"`
}
