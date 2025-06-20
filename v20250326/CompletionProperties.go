package v20250326

type CompletionProperties struct {
	HasMore Cursor `json:"hasMore"`
	Total   Cursor `json:"total"`
	Values  Values `json:"values"`
}
