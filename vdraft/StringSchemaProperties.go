package vdraft

type StringSchemaProperties struct {
	Default     Default      `json:"default"`
	Description Default      `json:"description"`
	Format      FormatClass  `json:"format"`
	MaxLength   Default      `json:"maxLength"`
	MinLength   Default      `json:"minLength"`
	Title       Default      `json:"title"`
	Type        JsonrpcClass `json:"type"`
}
