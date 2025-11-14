package vdraft

type ToolUseContentProperties struct {
	Meta  GetTaskPayloadResult `json:"_meta"`
	ID    Cursor               `json:"id"`
	Input Metadata             `json:"input"`
	Name  Cursor               `json:"name"`
	Type  JsonrpcClass         `json:"type"`
}
