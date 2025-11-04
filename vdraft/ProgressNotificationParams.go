package vdraft

type ProgressNotificationParams struct {
	Description string                               `json:"description"`
	Properties  ProgressNotificationParamsProperties `json:"properties"`
	Required    []string                             `json:"required"`
	Type        AnnotationsType                      `json:"type"`
}
