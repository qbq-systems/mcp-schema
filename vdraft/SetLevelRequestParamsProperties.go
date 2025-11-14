package vdraft

type SetLevelRequestParamsProperties struct {
	Meta  GetTaskPayloadResult `json:"_meta"`
	Level AnnotationsElement   `json:"level"`
	Task  AnnotationsElement   `json:"task"`
}
