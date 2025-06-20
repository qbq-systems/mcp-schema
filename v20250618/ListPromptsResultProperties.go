package v20250618

type ListPromptsResultProperties struct {
	Meta       Meta     `json:"_meta"`
	NextCursor Cursor   `json:"nextCursor"`
	Prompts    Messages `json:"prompts"`
}
