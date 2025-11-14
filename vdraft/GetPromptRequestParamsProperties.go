package vdraft

type GetPromptRequestParamsProperties struct {
	Meta      GetTaskPayloadResult `json:"_meta"`
	Arguments Arguments            `json:"arguments"`
	Name      Cursor               `json:"name"`
	Task      AnnotationsElement   `json:"task"`
}
