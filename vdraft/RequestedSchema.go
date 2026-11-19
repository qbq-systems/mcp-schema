package vdraft

type RequestedSchema struct {
	Description string                    `json:"description"`
	Properties  RequestedSchemaProperties `json:"properties"`
	Required    []string                  `json:"required"`
	Type        TypeElement               `json:"type"`
}
