package vdraft

type CompletionProperties struct {
	HasMore Cursor            `json:"hasMore"`
	Total   Cursor            `json:"total"`
	Values  SupportedVersions `json:"values"`
}
