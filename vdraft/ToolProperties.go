package vdraft

type ToolProperties struct {
	Meta         Meta             `json:"_meta"`
	Annotations  AnnotationsClass `json:"annotations"`
	Description  Cursor           `json:"description"`
	Icons        Audience         `json:"icons"`
	InputSchema  PutSchema        `json:"inputSchema"`
	Name         Cursor           `json:"name"`
	OutputSchema PutSchema        `json:"outputSchema"`
	Title        Cursor           `json:"title"`
}
