package vdraft

type ToolResultContentProperties struct {
	Meta              ClientResult      `json:"_meta"`
	Content           Audience          `json:"content"`
	IsError           Cursor            `json:"isError"`
	StructuredContent StructuredContent `json:"structuredContent"`
	ToolUseID         Cursor            `json:"toolUseId"`
	Type              JsonrpcClass      `json:"type"`
}
