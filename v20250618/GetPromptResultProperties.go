package v20250618

type GetPromptResultProperties struct {
	Meta        Meta     `json:"_meta"`
	Description Cursor   `json:"description"`
	Messages    Messages `json:"messages"`
}
