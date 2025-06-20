package v20250618

type CompletionProperties struct {
	HasMore Cursor `json:"hasMore"`
	Total   Cursor `json:"total"`
	Values  Values `json:"values"`
}
