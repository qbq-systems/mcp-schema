package vdraft

type CallToolRequestParamsProperties struct {
	Meta      GetTaskPayloadResult `json:"_meta"`
	Arguments GetTaskPayloadResult `json:"arguments"`
	Name      Cursor               `json:"name"`
	Task      AnnotationsElement   `json:"task"`
}
