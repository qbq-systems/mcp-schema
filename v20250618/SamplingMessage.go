package v20250618

type SamplingMessage struct {
	Description string                    `json:"description"`
	Properties  SamplingMessageProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
