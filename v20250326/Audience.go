package v20250326

type Audience struct {
	Description string      `json:"description"`
	Items       EmptyResult `json:"items"`
	Type        string      `json:"type"`
}
