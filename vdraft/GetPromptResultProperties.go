package vdraft

type GetPromptResultProperties struct {
	Meta        GetTaskPayloadResult `json:"_meta"`
	Description Cursor               `json:"description"`
	Messages    Messages             `json:"messages"`
}
