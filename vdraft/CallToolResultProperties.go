package vdraft

type CallToolResultProperties struct {
	Meta              Meta     `json:"_meta"`
	Content           Audience `json:"content"`
	IsError           Cursor   `json:"isError"`
	StructuredContent Meta     `json:"structuredContent"`
}
