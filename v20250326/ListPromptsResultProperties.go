package v20250326

type ListPromptsResultProperties struct {
	Meta       PurpleMeta `json:"_meta"`
	NextCursor Cursor     `json:"nextCursor"`
	Prompts    Messages   `json:"prompts"`
}
