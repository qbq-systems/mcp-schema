package vdraft

type InputSchemaProperties struct {
	Schema     Default       `json:"$schema"`
	Properties Properties3   `json:"properties"`
	Required   StopSequences `json:"required"`
	Type       JsonrpcClass  `json:"type"`
}
