package v20251125

type TitledSingleSelectEnumSchemaProperties struct {
	Default     Cursor       `json:"default"`
	Description Cursor       `json:"description"`
	OneOf       *Of          `json:"oneOf,omitempty"`
	Title       Cursor       `json:"title"`
	Type        JsonrpcClass `json:"type"`
	Enum        *Sizes       `json:"enum,omitempty"`
}
