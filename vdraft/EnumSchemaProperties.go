package vdraft

type EnumSchemaProperties struct {
	Default     Default       `json:"default"`
	Description Default       `json:"description"`
	Enum        StopSequences `json:"enum"`
	EnumNames   StopSequences `json:"enumNames"`
	Title       Default       `json:"title"`
	Type        JsonrpcClass  `json:"type"`
}
