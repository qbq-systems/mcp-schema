package vdraft

type CancelledNotification struct {
	Description string                          `json:"description"`
	Properties  CancelledNotificationProperties `json:"properties"`
	Required    []CallToolRequestRequired       `json:"required"`
	Type        AnnotationsType                 `json:"type"`
}
