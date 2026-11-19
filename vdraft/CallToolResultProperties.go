package vdraft

type CallToolResultProperties struct {
	Meta              AdditionalProperties `json:"_meta"`
	Content           Audience             `json:"content"`
	IsError           Cursor               `json:"isError"`
	ResultType        Cursor               `json:"resultType"`
	StructuredContent StructuredContent    `json:"structuredContent"`
}
