package v20251125

type CompletionProperties struct {
	HasMore Cursor `json:"hasMore"`
	Total   Cursor `json:"total"`
	Values  Sizes  `json:"values"`
}
