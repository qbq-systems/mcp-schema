package v20251125

type ListPromptsResultProperties struct {
	Meta       GetTaskPayloadResult `json:"_meta"`
	NextCursor Cursor               `json:"nextCursor"`
	Prompts    Messages             `json:"prompts"`
}
