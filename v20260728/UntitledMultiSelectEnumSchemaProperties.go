package v20260728

type UntitledMultiSelectEnumSchemaProperties struct {
	Default     SupportedVersions `json:"default"`
	Description Cursor            `json:"description"`
	Items       FluffyItems       `json:"items"`
	MaxItems    Cursor            `json:"maxItems"`
	MinItems    Cursor            `json:"minItems"`
	Title       Cursor            `json:"title"`
	Type        JsonrpcClass      `json:"type"`
}
