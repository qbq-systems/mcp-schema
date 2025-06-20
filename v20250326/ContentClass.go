package v20250326

type ContentClass struct {
	Items ClientNotification `json:"items"`
	Type  string             `json:"type"`
}
