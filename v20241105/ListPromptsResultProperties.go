package v20241105

type ListPromptsResultProperties struct {
	Meta       PurpleMeta `json:"_meta"`
	NextCursor Cursor     `json:"nextCursor"`
	Prompts    Messages   `json:"prompts"`
}
