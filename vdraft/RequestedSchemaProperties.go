package vdraft

type RequestedSchemaProperties struct {
	Properties FluffyProperties `json:"properties"`
	Required   StopSequences    `json:"required"`
	Type       JsonrpcClass     `json:"type"`
}
