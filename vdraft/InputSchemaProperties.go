package vdraft

type InputSchemaProperties struct {
	Properties TentacledProperties `json:"properties"`
	Required   StopSequences       `json:"required"`
	Type       JsonrpcClass        `json:"type"`
}
