package vdraft

type ResourceUpdatedNotificationParams struct {
	Description string                                      `json:"description"`
	Properties  ResourceUpdatedNotificationParamsProperties `json:"properties"`
	Required    []FormatElement                             `json:"required"`
	Type        AnnotationsType                             `json:"type"`
}
