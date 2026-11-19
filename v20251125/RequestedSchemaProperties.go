package v20251125

type RequestedSchemaProperties struct {
	Schema     Default           `json:"$schema"`
	Properties CunningProperties `json:"properties"`
	Required   StopSequences     `json:"required"`
	Type       JsonrpcClass      `json:"type"`
}
