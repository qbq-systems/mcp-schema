package v20250618

type EnumSchemaProperties struct {
	Description Default     `json:"description"`
	Enum        Enum        `json:"enum"`
	EnumNames   Enum        `json:"enumNames"`
	Title       Default     `json:"title"`
	Type        MethodClass `json:"type"`
}
