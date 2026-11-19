package v20251125

type CallToolRequestParamsProperties struct {
	Meta      GetTaskPayloadResult `json:"_meta"`
	Arguments GetTaskPayloadResult `json:"arguments"`
	Name      Cursor               `json:"name"`
	Task      AnnotationsElement   `json:"task"`
}
