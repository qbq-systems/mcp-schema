package vdraft

type InputSchemaProperties struct {
	Properties IndigoProperties `json:"properties"`
	Required   StopSequences    `json:"required"`
	Type       JsonrpcClass     `json:"type"`
}
