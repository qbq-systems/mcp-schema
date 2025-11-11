package vdraft

type LegacyTitledEnumSchemaProperties struct {
	Default     Default       `json:"default"`
	Description Default       `json:"description"`
	Enum        StopSequences `json:"enum"`
	EnumNames   Sizes         `json:"enumNames"`
	Title       Default       `json:"title"`
	Type        JsonrpcClass  `json:"type"`
}
