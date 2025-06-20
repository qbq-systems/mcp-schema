package v20250326

type ToolAnnotationsProperties struct {
	DestructiveHint Cursor `json:"destructiveHint"`
	IdempotentHint  Cursor `json:"idempotentHint"`
	OpenWorldHint   Cursor `json:"openWorldHint"`
	ReadOnlyHint    Cursor `json:"readOnlyHint"`
	Title           Cursor `json:"title"`
}
