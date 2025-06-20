package v20250618

type ToolProperties struct {
	Meta         Meta             `json:"_meta"`
	Annotations  AnnotationsClass `json:"annotations"`
	Description  Cursor           `json:"description"`
	InputSchema  PutSchema        `json:"inputSchema"`
	Name         Cursor           `json:"name"`
	OutputSchema PutSchema        `json:"outputSchema"`
	Title        Cursor           `json:"title"`
}
