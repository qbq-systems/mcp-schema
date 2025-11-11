package vdraft

type TitledMultiSelectEnumSchemaProperties struct {
	Default     Sizes        `json:"default"`
	Description Cursor       `json:"description"`
	Items       PurpleItems  `json:"items"`
	MaxItems    Cursor       `json:"maxItems"`
	MinItems    Cursor       `json:"minItems"`
	Title       Cursor       `json:"title"`
	Type        JsonrpcClass `json:"type"`
}
