package vdraft

type RequestedSchemaProperties struct {
	Properties HilariousProperties `json:"properties"`
	Required   Enum                `json:"required"`
	Type       MethodClass         `json:"type"`
}
