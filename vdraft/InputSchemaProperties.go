package vdraft

type InputSchemaProperties struct {
	Properties Properties3  `json:"properties"`
	Required   Enum         `json:"required"`
	Type       JsonrpcClass `json:"type"`
}
