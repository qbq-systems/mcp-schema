package vdraft

type CompleteRequestParamsProperties struct {
	Meta     GetTaskPayloadResult `json:"_meta"`
	Argument Argument             `json:"argument"`
	Context  Context              `json:"context"`
	Ref      ClientNotification   `json:"ref"`
	Task     AnnotationsElement   `json:"task"`
}
