package v20250618

type RequestedSchema struct {
	Description string                    `json:"description"`
	Properties  RequestedSchemaProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        AnnotationsType           `json:"type"`
}
