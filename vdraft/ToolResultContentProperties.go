package vdraft

type ToolResultContentProperties struct {
	Meta              Meta         `json:"_meta"`
	Content           Audience     `json:"content"`
	IsError           Cursor       `json:"isError"`
	StructuredContent Metadata     `json:"structuredContent"`
	ToolUseID         Cursor       `json:"toolUseId"`
	Type              JsonrpcClass `json:"type"`
}
