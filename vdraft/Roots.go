package vdraft

type Roots struct {
	Description string                    `json:"description"`
	Properties  AdditionalPropertiesClass `json:"properties"`
	Type        TypeElement               `json:"type"`
}
