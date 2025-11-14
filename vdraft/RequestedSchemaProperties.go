package vdraft

type RequestedSchemaProperties struct {
	Properties CunningProperties `json:"properties"`
	Required   StopSequences     `json:"required"`
	Type       JsonrpcClass      `json:"type"`
}
