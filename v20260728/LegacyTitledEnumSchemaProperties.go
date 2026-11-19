package v20260728

type LegacyTitledEnumSchemaProperties struct {
	Default     Default           `json:"default"`
	Description Default           `json:"description"`
	Enum        StopSequences     `json:"enum"`
	EnumNames   SupportedVersions `json:"enumNames"`
	Title       Default           `json:"title"`
	Type        JsonrpcClass      `json:"type"`
}
