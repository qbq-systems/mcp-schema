package v20241105

type CompletionProperties struct {
	HasMore Cursor `json:"hasMore"`
	Total   Cursor `json:"total"`
	Values  Values `json:"values"`
}
