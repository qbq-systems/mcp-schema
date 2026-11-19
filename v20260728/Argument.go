package v20260728

type Argument struct {
	Description string             `json:"description"`
	Properties  ArgumentProperties `json:"properties"`
	Required    []string           `json:"required"`
	Type        TypeElement        `json:"type"`
}
