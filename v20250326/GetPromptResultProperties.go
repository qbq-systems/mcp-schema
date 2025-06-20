package v20250326

type GetPromptResultProperties struct {
	Meta        PurpleMeta `json:"_meta"`
	Description Cursor     `json:"description"`
	Messages    Messages   `json:"messages"`
}
