package vdraft

type ToolResultContentProperties struct {
	Meta              GetTaskPayloadResult `json:"_meta"`
	Content           Audience             `json:"content"`
	IsError           Cursor               `json:"isError"`
	StructuredContent GetTaskPayloadResult `json:"structuredContent"`
	ToolUseID         Cursor               `json:"toolUseId"`
	Type              JsonrpcClass         `json:"type"`
}
