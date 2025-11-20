package vdraft

type ToolProperties struct {
	Meta         GetTaskPayloadResult `json:"_meta"`
	Annotations  AnnotationsElement   `json:"annotations"`
	Description  Cursor               `json:"description"`
	Execution    AnnotationsElement   `json:"execution"`
	Icons        Audience             `json:"icons"`
	InputSchema  PutSchema            `json:"inputSchema"`
	Name         Cursor               `json:"name"`
	OutputSchema PutSchema            `json:"outputSchema"`
	Title        Cursor               `json:"title"`
}
