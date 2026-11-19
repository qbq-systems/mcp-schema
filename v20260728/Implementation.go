package v20260728

type Implementation struct {
	Description string                   `json:"description"`
	Properties  ImplementationProperties `json:"properties"`
	Required    []string                 `json:"required"`
	Type        TypeElement              `json:"type"`
}
