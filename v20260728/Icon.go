package v20260728

type Icon struct {
	Description string         `json:"description"`
	Properties  IconProperties `json:"properties"`
	Required    []string       `json:"required"`
	Type        TypeElement    `json:"type"`
}
