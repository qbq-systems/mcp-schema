package vdraft

type ListPromptsResultProperties struct {
	Meta       GetTaskPayloadResult `json:"_meta"`
	NextCursor Cursor               `json:"nextCursor"`
	Prompts    Messages             `json:"prompts"`
}
