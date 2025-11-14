package vdraft

type RequestParamsProperties struct {
	Meta GetTaskPayloadResult `json:"_meta"`
	Task AnnotationsElement   `json:"task"`
}
