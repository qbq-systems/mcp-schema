package vdraft

type CancelledNotificationParams struct {
	Description string                                `json:"description"`
	Properties  CancelledNotificationParamsProperties `json:"properties"`
	Type        AnnotationsType                       `json:"type"`
}
