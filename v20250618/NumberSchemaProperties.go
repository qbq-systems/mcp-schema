package v20250618

type NumberSchemaProperties struct {
	Description Default     `json:"description"`
	Maximum     Default     `json:"maximum"`
	Minimum     Default     `json:"minimum"`
	Title       Default     `json:"title"`
	Type        FormatClass `json:"type"`
}
