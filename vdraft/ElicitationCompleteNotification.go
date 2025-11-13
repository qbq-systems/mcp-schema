package vdraft

type ElicitationCompleteNotification struct {
	Description string                                    `json:"description"`
	Properties  ElicitationCompleteNotificationProperties `json:"properties"`
	Required    []CallToolRequestRequired                 `json:"required"`
	Type        AnnotationsType                           `json:"type"`
}
