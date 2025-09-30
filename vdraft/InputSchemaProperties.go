package vdraft

type InputSchemaProperties struct {
	Properties Properties2  `json:"properties"`
	Required   Enum         `json:"required"`
	Type       JsonrpcClass `json:"type"`
}
