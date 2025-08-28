package vdraft

type LoggingMessageNotification struct {
	Description string                               `json:"description"`
	Properties  LoggingMessageNotificationProperties `json:"properties"`
	Required    []CallToolRequestRequired            `json:"required"`
	Type        AnnotationsType                      `json:"type"`
}
