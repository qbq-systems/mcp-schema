package vdraft

type ToolUseContentProperties struct {
	Meta  ClientResult      `json:"_meta"`
	ID    Cursor            `json:"id"`
	Input OutputSchemaClass `json:"input"`
	Name  Cursor            `json:"name"`
	Type  JsonrpcClass      `json:"type"`
}
