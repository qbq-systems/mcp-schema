package vdraft

type ProgressNotification struct {
	Description string                         `json:"description"`
	Properties  ProgressNotificationProperties `json:"properties"`
	Required    []CallToolRequestRequired      `json:"required"`
	Type        AnnotationsType                `json:"type"`
}
