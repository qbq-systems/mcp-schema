package v20251125

type CallToolResultProperties struct {
	Meta              GetTaskPayloadResult `json:"_meta"`
	Content           Audience             `json:"content"`
	IsError           Cursor               `json:"isError"`
	StructuredContent GetTaskPayloadResult `json:"structuredContent"`
}
