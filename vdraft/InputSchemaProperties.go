package vdraft

type InputSchemaProperties struct {
	Schema     Default          `json:"$schema"`
	Properties FriskyProperties `json:"properties"`
	Required   StopSequences    `json:"required"`
	Type       JsonrpcClass     `json:"type"`
}
