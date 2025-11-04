package vdraft

type CancelledNotificationParams struct {
	Description string                                `json:"description"`
	Properties  CancelledNotificationParamsProperties `json:"properties"`
	Required    []string                              `json:"required"`
	Type        AnnotationsType                       `json:"type"`
}
