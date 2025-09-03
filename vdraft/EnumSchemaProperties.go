package vdraft

type EnumSchemaProperties struct {
	Default     Default      `json:"default"`
	Description Default      `json:"description"`
	Enum        Enum         `json:"enum"`
	EnumNames   Enum         `json:"enumNames"`
	Title       Default      `json:"title"`
	Type        JsonrpcClass `json:"type"`
}
