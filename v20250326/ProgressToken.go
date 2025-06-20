package v20250326

type ProgressToken struct {
	Description string        `json:"description"`
	Type        []TypeElement `json:"type"`
}
