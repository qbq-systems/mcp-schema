package vdraft

type InputSchemaProperties struct {
	Properties FriskyProperties `json:"properties"`
	Required   StopSequences    `json:"required"`
	Type       JsonrpcClass     `json:"type"`
}
