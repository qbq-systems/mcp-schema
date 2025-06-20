package v20250618

type StringSchemaProperties struct {
	Description Default     `json:"description"`
	Format      FormatClass `json:"format"`
	MaxLength   Default     `json:"maxLength"`
	MinLength   Default     `json:"minLength"`
	Title       Default     `json:"title"`
	Type        MethodClass `json:"type"`
}
