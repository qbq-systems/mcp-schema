package vdraft

type RequestedSchemaProperties struct {
	Properties TentacledProperties `json:"properties"`
	Required   StopSequences       `json:"required"`
	Type       JsonrpcClass        `json:"type"`
}
