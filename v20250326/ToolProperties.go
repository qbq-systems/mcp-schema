package v20250326

type ToolProperties struct {
	Annotations AnnotationsClass `json:"annotations"`
	Description Cursor           `json:"description"`
	InputSchema InputSchema      `json:"inputSchema"`
	Name        Cursor           `json:"name"`
}
