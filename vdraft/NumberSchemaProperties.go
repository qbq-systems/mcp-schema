package vdraft

type NumberSchemaProperties struct {
	Default     Default     `json:"default"`
	Description Default     `json:"description"`
	Maximum     Default     `json:"maximum"`
	Minimum     Default     `json:"minimum"`
	Title       Default     `json:"title"`
	Type        FormatClass `json:"type"`
}
