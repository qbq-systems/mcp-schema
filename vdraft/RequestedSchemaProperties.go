package vdraft

type RequestedSchemaProperties struct {
	Schema     Default       `json:"$schema"`
	Properties JSONObject    `json:"properties"`
	Required   StopSequences `json:"required"`
	Type       JsonrpcClass  `json:"type"`
}
