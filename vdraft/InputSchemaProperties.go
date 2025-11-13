package vdraft

type InputSchemaProperties struct {
	Properties StickyProperties `json:"properties"`
	Required   StopSequences    `json:"required"`
	Type       JsonrpcClass     `json:"type"`
}
