package vdraft

type TaskAugmentedRequestParamsProperties struct {
	Meta GetTaskPayloadResult `json:"_meta"`
	Task AnnotationsElement   `json:"task"`
}
