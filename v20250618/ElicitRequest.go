package v20250618

type ElicitRequest struct {
	Description string                  `json:"description"`
	Properties  ElicitRequestProperties `json:"properties"`
	Required    []string                `json:"required"`
	Type        AnnotationsType         `json:"type"`
}
