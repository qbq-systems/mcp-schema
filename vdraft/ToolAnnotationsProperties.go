package vdraft

type ToolAnnotationsProperties struct {
	DestructiveHint Cursor       `json:"destructiveHint"`
	IdempotentHint  Cursor       `json:"idempotentHint"`
	OpenWorldHint   Cursor       `json:"openWorldHint"`
	ReadOnlyHint    Cursor       `json:"readOnlyHint"`
	TaskHint        LoggingLevel `json:"taskHint"`
	Title           Cursor       `json:"title"`
}
